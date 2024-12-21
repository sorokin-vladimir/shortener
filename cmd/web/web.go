package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

func StartServer(ctx context.Context, rdb *redis.Client) error {
	log.SetPrefix("WEB | ")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Example working with Redis
		ctx := r.Context()
		err := rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			http.Error(w, "Error working with Redis", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Welcome to web interface!")
	})

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
