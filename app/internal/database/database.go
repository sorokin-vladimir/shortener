package database

import (
	"context"
	"log"
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

	if redisUrl == "" {
		address := host + ":" + port
		return redis.NewClient(&redis.Options{
			Addr: address,
			DB:   dbNum,
		})
	}

	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatalf("Could not parse Redis URL: %v", err)
	}
	return redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Username: opts.Username,
		Password: opts.Password,
		DB:       dbNum,
	})
}
