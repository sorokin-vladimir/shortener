package shortener

import (
	"testing"
)

func TestBase56Encode(t *testing.T) {
	tests := []struct {
		name     string
		number   uint64
		expected string
		_ 			 string
	}{
		{name: "encode zero", number: 0, expected: ""},
		{name: "encode 1", number: 1, expected: "b"},
		{name: "encode 998", number: 998, expected: "1u"},
		{name: "encode 999", number: 999, expected: "2u"},
		{name: "encode 15,000,000", number: 15000000, expected: "klAGb"},
		{name: "encode 100,000,000", number: 100000000, expected: "tWAlm"},
		{name: "encode biggest number", number: 18446744073709551615, expected: "sDC7KC3C3Zeb"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Base56Encode(tt.number)
			if result != tt.expected {
					t.Errorf("Base56Encode(%d) = %s; expected %s", tt.number, result, tt.expected)
			}
	})
	}
}
