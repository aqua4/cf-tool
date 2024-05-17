package main

import (
	"fmt"
	"os"
	"strings"

	"cf-tool/client"
	"cf-tool/cmd"
	"cf-tool/config"
	"github.com/mitchellh/go-homedir"

	docopt "github.com/docopt/docopt-go"
)

const version = "v1.0.11"
const configPath = "~/.cf/config"
const sessionPath = "~/.cf/session"

func main() {
	usage := `Codeforces Tool $%version%$ (cf). https://github.com/aqua4/cf-tool
You should run "cf config" to configure your handle and password at first.
Usage:
  cf config
  cf submit -f <file> <url>
Options:
  -h --help     Show this screen.
  --version     Show version.
  -f <file>, --file <file>
                Path to file. E.g. "a.cpp", "./temp/a.cpp"
  <url>         Problem URL. E.g. "https://codeforces.com/contest/180/problem/A",
Examples:
  cf config            Configure the cf-tool.
  cf submit -f a.cpp https://codeforces.com/contest/100/A
File:
  cf will save some data in some files:
  "~/.cf/config"        Configuration file.
  "~/.cf/session"       Session file, including cookies, handle, password, etc.
  "~" is the home directory of current user in your system.`
	usage = strings.Replace(usage, `$%version%$`, version, 1)
	opts, _ := docopt.ParseArgs(usage, os.Args[1:], fmt.Sprintf("Codeforces Tool (cf) %v\n", version))
	opts[`{version}`] = version
	cfgPath, _ := homedir.Expand(configPath)
	clnPath, _ := homedir.Expand(sessionPath)
	config.Init(cfgPath)
	client.Init(clnPath, config.Instance.Host, config.Instance.Proxy)
	err := cmd.Eval(opts)
	if err != nil {
		fmt.Println(err.Error())
	}
}
