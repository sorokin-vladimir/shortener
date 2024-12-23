package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

const (
	DB_SHORT  = 0
	DB_LIMITS = 1
)

var Ctx = context.Background()

func CreateClient(dbNum int) *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	redisUrl := os.Getenv("REDIS_URL")

	var address string
	if redisUrl == "" {
		address = host + ":" + port
	} else {
		address = redisUrl
	}

	return redis.NewClient(&redis.Options{
		Addr: address,
		DB:   dbNum,
	})
}
