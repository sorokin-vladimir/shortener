package shortener

import (
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/sorokin-vladimir/shortener/internal/database"
)

func encode(number uint64, db_shorts *redis.Client) string {
	var encodedBuilder strings.Builder
	length := len(alphabet)

	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
		tmpStr := encodedBuilder.String()
		if len(tmpStr) > 3 {
			val, _ := db_shorts.Get(database.Ctx, tmpStr).Result()
			if val == "" {
				return tmpStr
			}
		}
	}

	return encodedBuilder.String()
}
