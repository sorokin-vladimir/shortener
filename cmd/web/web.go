package web

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

func StartServer(ctx context.Context, rdb *redis.Client) error {
	log.SetPrefix("WEB | ")
	http.HandleFunc("/{url}", resoleShort)

	port := os.Getenv("WEB_PORT")
	server := &http.Server{Addr: ":" + port}

	// Run the server
	go func() {
		<-ctx.Done()
		log.Println("Stopping web server...")
		server.Shutdown(ctx)
	}()

	log.Printf("Web server is run on port: %s", port)
	return server.ListenAndServe()
}
