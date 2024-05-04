package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Config load and save configuration
type Config struct {
	Host  string `json:"host"`
	Proxy string `json:"proxy"`
	path  string
}

// Instance global configuration
var Instance *Config

// Init initialize
func Init(path string) {
	c := &Config{path: path, Host: "https://codeforces.com", Proxy: ""}
	if err := c.load(); err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Create a new configuration in %v\n", path)
	}
	if err := c.save(); err != nil {
		fmt.Println(err)
	}
	Instance = c
}

// load from path
func (c *Config) load() (err error) {
	file, err := os.Open(c.path)
	if err != nil {
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}

// save file to path
func (c *Config) save() (err error) {
	var data bytes.Buffer
	encoder := json.NewEncoder(&data)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(c)
	if err == nil {
		err = os.MkdirAll(filepath.Dir(c.path), os.ModePerm)
		if err == nil {
			err = os.WriteFile(c.path, data.Bytes(), 0644)
		}
	} else {
		err = fmt.Errorf("Cannot save config to %v\n%v", c.path, err.Error())
	}
	return
}
