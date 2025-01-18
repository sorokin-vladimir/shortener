package web

import (
	"context"
	"log"
	"net/http"
	"os"
)

func StartServer(ctx context.Context) error {
	// http.HandleFunc("/{url}", resolveShort)
	http.HandleFunc("/health", health)
	// http.HandleFunc("/short", short)

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
