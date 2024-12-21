package utils

import (
	"net/url"
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func IsLocalhost(url string) bool {
	lowerUrl := strings.ToLower(url)
	domain := os.Getenv("DOMAIN")

	if lowerUrl == domain {
		return false
	}

	if strings.Contains(lowerUrl, "localhost:") || strings.Contains(lowerUrl, "127.0.0.1") {
		return true
	}

	return false
}

func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
