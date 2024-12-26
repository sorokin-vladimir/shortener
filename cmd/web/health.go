package web

import (
	"log"
	"net/http"

	"github.com/sorokin-vladimir/shortener/internal/database"
)

func health(w http.ResponseWriter, _ *http.Request) {
	db_shorts := database.CreateClient(0)
	defer db_shorts.Close()

	keys, err := db_shorts.Keys(database.Ctx, "*").Result()

	if err != nil {
		log.Println("Health check: Internal error", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	if len(keys) >= 0 {
		http.Error(w, "OK", http.StatusOK)
		return
	}
}
