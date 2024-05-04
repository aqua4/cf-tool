package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"cf-tool/cookiejar"
)

// Client codeforces client
type Client struct {
	Jar           *cookiejar.Jar `json:"cookies"`
	Handle        string         `json:"handle"`
	HandleOrEmail string         `json:"handle_or_email"`
	Password      string         `json:"password"`
	Ftaa          string         `json:"ftaa"`
	Bfaa          string         `json:"bfaa"`
	host          string
	proxy         string
	path          string
	client        *http.Client
}

// Instance global client
var Instance *Client

// Init initialize
func Init(path, host, proxy string) {
	jar, _ := cookiejar.New(nil)
	c := &Client{Jar: jar, path: path, host: host, proxy: proxy, client: nil}
	if err := c.load(); err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Create a new session in %v\n", path)
	}
	Proxy := http.ProxyFromEnvironment
	if len(proxy) > 0 {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Use default proxy from environment")
		} else {
			Proxy = http.ProxyURL(proxyURL)
		}
	}
	c.client = &http.Client{Jar: c.Jar, Transport: &http.Transport{Proxy: Proxy}}
	if err := c.save(); err != nil {
		fmt.Println(err.Error())
	}
	Instance = c
}

// load from path
func (c *Client) load() (err error) {
	file, err := os.Open(c.path)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, c)
}

// save file to path
func (c *Client) save() (err error) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err == nil {
		err = os.MkdirAll(filepath.Dir(c.path), os.ModePerm)
		if err == nil {
			err = os.WriteFile(c.path, data, 0644)
		}
	}
	if err != nil {
		fmt.Printf("Cannot save session to %v\n%v\n", c.path, err.Error())
	}
	return
}
