package cmd

import (
	"cf-tool/client"
	"cf-tool/config"
	"cf-tool/util"
	"fmt"
)

// Config command
func Config() (err error) {
	cfg := config.Instance
	cln := client.Instance
	fmt.Println("Configure the tool")
	fmt.Println(`0) login`)
	fmt.Println(`1) set host domain`)
	fmt.Println(`2) set proxy`)
	index := util.ChooseIndex(3)
	if index == 0 {
		return cln.ConfigLogin()
	} else if index == 1 {
		return cfg.SetHost()
	} else if index == 2 {
		return cfg.SetProxy()
	}
	return
}
