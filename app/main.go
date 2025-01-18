package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/sorokin-vladimir/shortener/cmd/telegram"
	"github.com/sorokin-vladimir/shortener/cmd/web"
	// "github.com/sorokin-vladimir/shortener/internal/database"
)

func main() {
	// Create a context for the management application
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// db_shorts := database.CreateClient(database.DB_SHORT)

	// // Check the Redis connection
	// if err := db_shorts.Ping(ctx).Err(); err != nil {
	// 	log.Fatalf("Could not connect to Redis: %v", err)
	// }
	// log.Println("Connection to Redis is completed!")

	// Run Web server
	go func() {
		if err := web.StartServer(ctx); err != nil {
			log.Fatalf("Error of the web-server running: %v", err)
		}
	}()

	// Run Telegram bot
	// go func() {
	// 	if err := telegram.StartBot(ctx); err != nil {
	// 		log.Fatalf("Error of the Telegram-bot running: %v", err)
	// 	}
	// }()

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
