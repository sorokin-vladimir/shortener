package main

import (
	"context"
	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/sorokin-vladimir/shortener/cmd/web"
	"github.com/sorokin-vladimir/shortener/cmd/telegram"
)

func main() {
	log.Println("main func")

	// Create a context for the management application
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		Password: "", // Password Redis, if exists
		DB:       0,
	})

	// Check the Redis connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connection to Redis is completed!")

	// Run Web server
	go func() {
		if err := web.StartServer(ctx, rdb); err != nil {
			log.Fatalf("Error of the web-server running: %v", err)
		}
	}()

	// Run Telegram bot
	go func() {
		if err := telegram.StartBot(ctx, rdb); err != nil {
			log.Fatalf("Error of the Telegram-bot running: %v", err)
		}
	}()

	// Waiting shutdown signal
	waitForShutdown(cancel)
}

// The function for waiting shutdown signal
func waitForShutdown(cancel context.CancelFunc) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("\nEnding work...")
	cancel()
	time.Sleep(2 * time.Second) // Waiting end of gorutins
}
