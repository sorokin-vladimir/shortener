package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
    log.Println("Telegram: Test running docker and connection to redis... qq")

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

    log.Println("Telegram: Connected to the database")

    log.Println("Telegram: ##########################")
}
