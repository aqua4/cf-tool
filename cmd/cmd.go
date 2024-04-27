package cmd

import (
	"fmt"

	"github.com/docopt/docopt-go"

	"cf-tool/client"
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
	}
	return nil
}

func loginAgain(cln *client.Client, err error) error {
	if err != nil && err.Error() == client.ErrorNotLogged {
		fmt.Println("Not logged. Try to login\n")
		err = cln.Login()
	}
	return err
}
