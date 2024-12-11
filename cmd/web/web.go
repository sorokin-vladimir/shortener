package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func StartServer(ctx context.Context, rdb *redis.Client) error {
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

	// TODO: move port num into envs
	server := &http.Server{Addr: ":8080"}

	// Run the server
	go func() {
		<-ctx.Done()
		log.Println("Stopping web server...")
		server.Shutdown(ctx)
	}()

	log.Println("Web server is run on port: 8080")
	return server.ListenAndServe()
}
