package shortener

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {

	for _, tt := range TestData {
		t.Run(fmt.Sprintf("Decode %s", tt.encodedStr), func(t *testing.T) {
			result, err := Decode(tt.encodedStr)

			if err != nil {
				t.Error(err)
			}

			if result != tt.number {
				t.Errorf("Decode(%s) = %d; expected %d", tt.encodedStr, result, tt.number)
			}
		})
	}

	const encodedStrWithWrongSymbol = "abc_e"
	t.Run(fmt.Sprintf("Decode %s", encodedStrWithWrongSymbol), func(t *testing.T) {
		result, err := Decode(encodedStrWithWrongSymbol)

		if err != nil && err.Error() == "wrong symbol" {
			t.Logf("Decode(%s): expected error", encodedStrWithWrongSymbol)
			return
		}

		t.Errorf("Decode(%s) = %d; expected error", encodedStrWithWrongSymbol, result)
	})
}
