package client

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"cf-tool/util"

	"github.com/PuerkitoBio/goquery"
)

// Submission submit state
type Submission struct {
	status string
	time   uint64
	memory uint64
	end    bool
}

func isWait(verdict string) bool {
	return verdict == "null" || verdict == "TESTING" || verdict == "SUBMITTED"
}

// GetStatus returns status
func (s *Submission) GetStatus() string {
	return s.status
}

// ParseMemory formatter
func (s *Submission) ParseMemory() string {
	if s.memory > 1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(s.memory)/1024.0/1024.0)
	} else if s.memory > 1024 {
		return fmt.Sprintf("%.2f KB", float64(s.memory)/1024.0)
	}
	return fmt.Sprintf("%v B", s.memory)
}

// ParseTime formatter
func (s *Submission) ParseTime() string {
	return fmt.Sprintf("%v ms", s.time)
}

func findSubmission(body []byte) ([]byte, error) {
	reg := regexp.MustCompile(`data-submission-id=['"]\d[\s\S]+?</tr>`)
	tmp := reg.Find(body)
	if tmp == nil {
		return nil, errors.New("Cannot find any submission")
	}
	return tmp, nil
}

const Running = "Running"

func parseSubmission(body []byte) (ret Submission, err error) {
	data := fmt.Sprintf("<table><tr %v</table>", string(body))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`\d+`)
	sub := doc.Find(".submissionVerdictWrapper")
	end := false
	if verdict, exist := sub.Attr("submissionverdict"); exist && !isWait(verdict) {
		end = true
	}
	status, _ := sub.Html()
	numReg := regexp.MustCompile(`\d+`)
	fmtReg := regexp.MustCompile(`<span\sclass=["']?verdict-format-([\S^>]+?)["']?>`)
	colReg := regexp.MustCompile(`<span\sclass=["']?verdict-([\S^>]+?)["']?>`)
	tagReg := regexp.MustCompile(`<[\s\S]*?>`)
	status = fmtReg.ReplaceAllString(status, "")
	status = colReg.ReplaceAllString(status, "")
	status = tagReg.ReplaceAllString(status, "")
	status = strings.TrimSpace(status)
	// empty status usually means that the submission is in queue
	if status == "" {
		status = Running
	}
	if s := numReg.FindString(status); s != "" {
		status = strings.ReplaceAll(status, "${f-points}", s)
		status = strings.ReplaceAll(status, "${f-passed}", s)
		status = strings.ReplaceAll(status, "${f-judged}", s)
	}
	getInt := func(sel string) uint64 {
		if tmp := reg.FindString(doc.Find(sel).Text()); tmp != "" {
			t, _ := strconv.Atoi(tmp)
			return uint64(t)
		}
		return 0
	}
	time := getInt(".time-consumed-cell")
	memory := getInt(".memory-consumed-cell") * 1024
	return Submission{
		status,
		time,
		memory,
		end,
	}, nil
}

func (c *Client) getSubmission(URL string) (submission Submission, err error) {
	body, err := util.GetBody(c.client, URL)
	if err != nil {
		return
	}
	submissionBody, err := findSubmission(body)
	if err != nil {
		return
	}
	return parseSubmission(submissionBody)
}

// WatchSubmission watches submission
func (c *Client) WatchSubmission(info Info) error {
	URL, err := info.MySubmissionURL(c.host)
	if err != nil {
		return err
	}

	for {
		st := time.Now()
		submission, err := c.getSubmission(URL)
		if err != nil {
			return err
		}
		status := submission.GetStatus()
		fmt.Printf("status: %v\n", status)
		if !strings.HasPrefix(status, Running) {
			fmt.Printf("time: %v\n", submission.ParseTime())
			fmt.Printf("memory: %v\n", submission.ParseMemory())
		}
		if submission.end {
			return nil
		}
		sub := time.Since(st)
		if sub < time.Second {
			time.Sleep(time.Duration(time.Second - sub))
		}
	}
}
