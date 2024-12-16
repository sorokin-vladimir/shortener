package shortener

import (
	"errors"
	"math"
	"strings"
)

func Decode(encodedString string) (uint64, error) {
	var number uint64
	length := len(Alphabet)

	for i, symbol := range encodedString {
		alphabeticPosition := strings.IndexRune(Alphabet, symbol)
		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("wrong symbol")
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
