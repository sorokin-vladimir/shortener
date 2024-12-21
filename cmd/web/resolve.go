package web

import (
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/sorokin-vladimir/shortener/internal/database"
)

func resoleShort(w http.ResponseWriter, r *http.Request) {
	db_shorts := database.CreateClient(0)
	defer db_shorts.Close()

	value, err := db_shorts.Get(database.Ctx, r.PathValue("url")).Result()

	if err == redis.Nil {
		log.Printf("Short-URL '%s' not found", r.PathValue("url"))
		http.Error(w, "short-url not found in db", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println("Internal error", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	db_limits := database.CreateClient(1)
	defer db_limits.Close()

	_ = db_limits.Incr(database.Ctx, "counter")

	http.Redirect(w, r, value, http.StatusSeeOther)
}
