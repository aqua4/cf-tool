package util

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// CHA map
const CHA = "abcdefghijklmnopqrstuvwxyz0123456789"

// RandString n is the length. a-z 0-9
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = CHA[rand.Intn(len(CHA))]
	}
	return string(b)
}

// Scanline scan line
func Scanline() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	fmt.Println("\nInterrupted.")
	os.Exit(1)
	return ""
}

// ScanlineTrim scan line and trim
func ScanlineTrim() string {
	return strings.TrimSpace(Scanline())
}

// ChooseIndex return valid index in [0, maxLen)
func ChooseIndex(maxLen int) int {
	fmt.Println("Please choose one (index): ")
	for {
		index := ScanlineTrim()
		i, err := strconv.Atoi(index)
		if err == nil && i >= 0 && i < maxLen {
			return i
		}
		fmt.Println("Invalid index! Please try again: ")
	}
}

// GetBody read body
func GetBody(client *http.Client, URL string) ([]byte, error) {
	resp, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// PostBody read post body
func PostBody(client *http.Client, URL string, data url.Values) ([]byte, error) {
	resp, err := client.PostForm(URL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
