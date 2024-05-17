package cmd

import (
	"fmt"
	"regexp"

	"cf-tool/client"

	"github.com/docopt/docopt-go"
)

// ParsedArgs parsed arguments
type ParsedArgs struct {
	Info    client.Info
	File    string
	Url     string `docopt:"<url>"`
	Version string `docopt:"{version}"`
	Config  bool   `docopt:"config"`
	Submit  bool   `docopt:"submit"`
}

// Args global variable
var Args *ParsedArgs

func parseArgs(opts docopt.Opts) error {
	if file, ok := opts["--file"].(string); ok {
		Args.File = file
	}
	info := client.Info{}
	parsed := parseArg(Args.Url)
	if value, ok := parsed["problemType"]; ok {
		if info.ProblemType != "" && info.ProblemType != value {
			return fmt.Errorf("Problem Type conflicts: %v %v", info.ProblemType, value)
		}
		info.ProblemType = value
	}
	if value, ok := parsed["contestID"]; ok {
		if info.ContestID != "" && info.ContestID != value {
			return fmt.Errorf("Contest ID conflicts: %v %v", info.ContestID, value)
		}
		info.ContestID = value
	}
	if value, ok := parsed["problemID"]; ok {
		if info.ProblemID != "" && info.ProblemID != value {
			return fmt.Errorf("Problem ID conflicts: %v %v", info.ProblemID, value)
		}
		if value == "0" {
			value = "A"
		}
		info.ProblemID = value
	}
	if info.ProblemType == "" || info.ProblemType == "contest" {
		if len(info.ContestID) < 6 {
			info.ProblemType = "contest"
		} else {
			info.ProblemType = "gym"
		}
	}
	Args.Info = info
	return nil
}

// ProblemRegStr problem
const ProblemRegStr = `\w+`

// ContestRegStr contest
const ContestRegStr = `\d+`

// ArgRegStr for parsing arg
var ArgRegStr = [...]string{
	fmt.Sprintf(`/contest/(?P<contestID>%v)(/problem/(?P<problemID>%v))?`, ContestRegStr, ProblemRegStr),
	fmt.Sprintf(`/gym/(?P<contestID>%v)(/problem/(?P<problemID>%v))?`, ContestRegStr, ProblemRegStr),
	fmt.Sprintf(`/problemset/problem/(?P<contestID>%v)/(?P<problemID>%v)`, ContestRegStr, ProblemRegStr),
}

// ArgType type
var ArgType = [...]string{
	"contest",
	"gym",
	"contest",
}

func parseArg(arg string) map[string]string {
	output := make(map[string]string)
	for k, regStr := range ArgRegStr {
		reg := regexp.MustCompile(regStr)
		names := reg.SubexpNames()
		for i, val := range reg.FindStringSubmatch(arg) {
			if names[i] != "" && val != "" {
				output[names[i]] = val
			}
			output["problemType"] = ArgType[k]
		}
	}
	return output
}
