package client

import (
	"errors"
	"fmt"
	"strings"
)

// Info information
type Info struct {
	ProblemType string `json:"problem_type"`
	ContestID   string `json:"contest_id"`
	ProblemID   string `json:"problem_id"`
}

// ErrorNeedProblemID error
const ErrorNeedProblemID = "You have to specify the Problem ID"

// ErrorNeedContestID error
const ErrorNeedContestID = "You have to specify the Contest ID"

// ErrorNeedGymID error
const ErrorNeedGymID = "You have to specify the Gym ID"

// ErrorUnknownType error
const ErrorUnknownType = "Unknown Type"

func (info *Info) errorContest() (string, error) {
	if info.ProblemType == "gym" {
		return "", errors.New(ErrorNeedGymID)
	}
	return "", errors.New(ErrorNeedContestID)
}

// Hint hint text
func (info *Info) Hint() string {
	text := strings.ToUpper(info.ProblemType)
	if info.ContestID != "" {
		text = text + " " + info.ContestID
	}
	if info.ProblemID != "" {
		text = text + ", problem " + info.ProblemID
	}
	return text
}

// ProblemSetURL parse problem set url
func (info *Info) ProblemSetURL(host string) (string, error) {
	if info.ContestID == "" {
		return info.errorContest()
	}
	switch info.ProblemType {
	case "contest":
		return fmt.Sprintf(host+"/contest/%v", info.ContestID), nil
	case "gym":
		return fmt.Sprintf(host+"/gym/%v", info.ContestID), nil
	}
	return "", errors.New(ErrorUnknownType)
}

// MySubmissionURL parse submission url
func (info *Info) MySubmissionURL(host string) (string, error) {
	if info.ContestID == "" {
		return info.errorContest()
	}
	switch info.ProblemType {
	case "contest":
		return fmt.Sprintf(host+"/contest/%v/my", info.ContestID), nil
	case "gym":
		return fmt.Sprintf(host+"/gym/%v/my", info.ContestID), nil
	}
	return "", errors.New(ErrorUnknownType)
}

// SubmitURL submit url
func (info *Info) SubmitURL(host string) (string, error) {
	URL, err := info.ProblemSetURL(host)
	if err != nil {
		return "", err
	}
	return URL + "/submit", nil
}
