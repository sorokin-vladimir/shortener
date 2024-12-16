package shortener

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, tt := range TestData {
		t.Run(fmt.Sprintf("Encode %d", tt.number), func(t *testing.T) {
			result := Encode(tt.number)
			if result != tt.encodedStr {
				t.Errorf("Encode(%d) = %s; expected %s", tt.number, result, tt.encodedStr)
			}
		})
	}
}
