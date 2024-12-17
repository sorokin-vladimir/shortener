package utils

import (
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
	if strings.Contains(lowerUrl, "localhost:") || strings.Contains(lowerUrl, "127.0.0.1") {
		return true
	}

	return false
}
