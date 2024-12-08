package shortener

import (
	"testing"
)

func TestBase56Decode(t *testing.T) {
	tests := []struct {
		name     			string
		encodedString string
		expected 			uint64
		expectedError string
		_ 						string
	}{
		{name: "decode empty string", encodedString: "", expected: 0},
		{name: "decode b", encodedString: "b", expected: 1},
		{name: "decode 1u", encodedString: "1u", expected: 998},
		{name: "decode 2u", encodedString: "2u", expected: 999},
		{name: "decode klAGb", encodedString: "klAGb", expected: 15000000},
		{name: "decode tWAlm", encodedString: "tWAlm", expected: 100000000},
		{name: "decode sDC7KC3C3Zeb", encodedString: "sDC7KC3C3Zeb", expected: 18446744073709551615},
		{name: "decode _ with error", encodedString: "_", expected: 0, expectedError: "wrong symbol"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Base56Decode(tt.encodedString)

			if err != nil && err.Error() == tt.expectedError {
				t.Logf("Base56Decode(%s): expected error", tt.encodedString)
				return
			}

			if err != nil && err.Error() != tt.expectedError {
				t.Error(err)
			}

			if result != tt.expected {
					t.Errorf("Base56Decode(%s) = %d; expected %d", tt.encodedString, result, tt.expected)
			}
	})
	}
}
