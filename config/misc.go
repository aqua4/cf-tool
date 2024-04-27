package config

import (
	"fmt"
	"regexp"

	"cf-tool/util"
)

func formatHost(host string) (string, error) {
	reg := regexp.MustCompile(`https?://[\w\-]+(\.[\w\-]+)+/?`)
	if !reg.MatchString(host) {
		return "", fmt.Errorf(`Invalid host "%v"`, host)
	}
	for host[len(host)-1:] == "/" {
		host = host[:len(host)-1]
	}
	return host, nil
}

func formatProxy(proxy string) (string, error) {
	reg := regexp.MustCompile(`[\w\-]+?://[\w\-]+(\.[\w\-]+)*(:\d+)?`)
	if !reg.MatchString(proxy) {
		return "", fmt.Errorf(`Invalid proxy "%v"`, proxy)
	}
	return proxy, nil
}

// SetHost set host for Codeforces
func (c *Config) SetHost() (err error) {
	host, err := formatHost(c.Host)
	if err != nil {
		host = "https://codeforces.com"
	}
	fmt.Println("Current host domain is %v", host)
	fmt.Println(`Set a new host domain (e.g. "https://codeforces.com"`)
	fmt.Println(`Note: Don't forget the "http://" or "https://"`)
	for {
		host, err = formatHost(util.ScanlineTrim())
		if err == nil {
			break
		}
		fmt.Println(err.Error())
	}
	c.Host = host
	fmt.Println("New host domain is %v", host)
	return c.save()
}

// SetProxy set proxy for client
func (c *Config) SetProxy() (err error) {
	proxy, err := formatProxy(c.Proxy)
	if err != nil {
		proxy = ""
	}
	if len(proxy) == 0 {
		fmt.Println("Current proxy is based on environment")
	} else {
		fmt.Println("Current proxy is %v", proxy)
	}
	fmt.Println(`Set a new proxy (e.g. "http://127.0.0.1:2333", "socks5://127.0.0.1:1080"`)
	fmt.Println(`Enter empty line if you want to use default proxy from environment`)
	fmt.Println(`Note: Proxy URL should match "protocol://host[:port]"`)
	for {
		proxy, err = formatProxy(util.ScanlineTrim())
		if err == nil {
			break
		}
		fmt.Println(err.Error())
	}
	c.Proxy = proxy
	if len(proxy) == 0 {
		fmt.Println("Current proxy is based on environment")
	} else {
		fmt.Println("Current proxy is %v", proxy)
	}
	return c.save()
}
