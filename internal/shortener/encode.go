package shortener

import (
	"strings"
)

func Base56Encode(number uint64) string {
	length := len(Alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(Alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}
