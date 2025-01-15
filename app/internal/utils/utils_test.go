package utils

import (
	"testing"
)

func TestEnforceHTTP(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"http://example.com/", "http://example.com/"},
		{"https://example.com/", "https://example.com/"},
		{"example.com/", "http://example.com/"},
	}

	for _, tt := range tests {
		t.Run("Check "+tt.url, func(t *testing.T) {
			result := EnforceHTTP(tt.url)
			if result != tt.expected {
				t.Errorf("EnforceHTTP(%s) = %s; expected %s", tt.url, result, tt.expected)
			}
		})
	}
}

func TestIsLocalhost(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"http://example.com/", false},
		{"https://example.com/", false},
		{"localhost:8080/qwerty", true},
		{"http://localhost:8080/qwerty", true},
		{"http://locAlhoSt:8080/qwerty", true},
		{"127.0.0.1:8080/qwerty", true},
	}

	for _, tt := range tests {
		t.Run("Is localhost: "+tt.url, func(t *testing.T) {
			result := IsLocalhost(tt.url)
			if result != tt.expected {
				t.Errorf("EnforceHTTP(%s) = %t; expected %t", tt.url, result, tt.expected)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"http://example.com/", true},
		{"https://example.com/", true},
		{"http://example.com", true},
		{"http://example", true},
		{"example.com", false},
		{"example.com/", false},
		{"example", false},
	}

	for _, tt := range tests {
		t.Run("Is URL: "+tt.str, func(t *testing.T) {
			result := IsURL(tt.str)
			if result != tt.expected {
				t.Errorf("IsURL(%s) = %t; expected %t", tt.str, result, tt.expected)
			}
		})
	}
}
