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
	password := os.Getenv("REDIS_PASSWORD")

	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       dbNum,
	})
}
