package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/docopt/docopt-go"

	"cf-tool/client"

	"github.com/fatih/color"
)

// Eval opts
func Eval(opts docopt.Opts) error {
	Args = &ParsedArgs{}
	err := opts.Bind(Args)
	if err != nil {
		return err
	}
	if err := parseArgs(opts); err != nil {
		return err
	}
	if Args.Config {
		return Config()
	} else if Args.Submit {
		return Submit()
	} else if Args.Upgrade {
		return Upgrade()
	}
	return nil
}

func getSampleID() (samples []string) {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	paths, err := os.ReadDir(path)
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`testI(\d+).txt`)
	for _, path := range paths {
		name := path.Name()
		tmp := reg.FindSubmatch([]byte(name))
		if tmp != nil {
			idx := string(tmp[1])
			ans := fmt.Sprintf("testO%v.txt", idx)
			if _, err := os.Stat(ans); err == nil {
				samples = append(samples, idx)
			}
		}
	}
	return
}

func loginAgain(cln *client.Client, err error) error {
	if err != nil && err.Error() == client.ErrorNotLogged {
		color.Red("Not logged. Try to login\n")
		err = cln.Login()
	}
	return err
}
