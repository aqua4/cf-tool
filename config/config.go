package config

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/color"
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
		color.Red(err.Error())
		color.Green("Create a new configuration in %v", path)
	}
	c.save()
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
		os.MkdirAll(filepath.Dir(c.path), os.ModePerm)
		err = os.WriteFile(c.path, data.Bytes(), 0644)
	}
	if err != nil {
		color.Red("Cannot save config to %v\n%v", c.path, err.Error())
	}
	return
}
