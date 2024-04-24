package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"cf-tool/client"
)

// Submit command
func Submit() (err error) {
	cln := client.Instance
	info := Args.Info
	filename := Args.File

	ext := filepath.Ext(filename)
	langID, ok := client.ExtToLangID[ext]
	if !ok {
		return fmt.Errorf("%v can not match any supported file extension.", ext)
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	source := string(bytes)

	if err = cln.Submit(info, langID, source); err != nil {
		if err = loginAgain(cln, err); err == nil {
			err = cln.Submit(info, langID, source)
		}
	}
	return
}
