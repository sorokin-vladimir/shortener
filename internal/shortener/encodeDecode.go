package shortener

import (
	"errors"
	"math"
	"strings"
)

func encode(number uint64) string {
	var encodedBuilder strings.Builder
	length := len(alphabet)

	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}

func decode(encodedString string) (uint64, error) {
	var number uint64
	length := len(alphabet)

	for i, symbol := range encodedString {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)
		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("wrong symbol")
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
