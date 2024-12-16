package shortener

import (
	"strings"
)

func Encode(number uint64) string {
	var encodedBuilder strings.Builder
	length := len(Alphabet)

	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(Alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}
