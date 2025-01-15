package web

import (
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/sorokin-vladimir/shortener/internal/database"
)

func resolveShort(w http.ResponseWriter, r *http.Request) {
	db_shorts := database.CreateClient(database.DB_SHORT)
	defer db_shorts.Close()

	url := r.PathValue("url")
	value, err := db_shorts.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		log.Printf("Could not resolve '%s'", url)
		http.Error(w, "Could not resolve "+url, http.StatusNotFound)
		return
	} else if err != nil {
		log.Println("Internal error", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	db_limits := database.CreateClient(database.DB_LIMITS)
	defer db_limits.Close()

	_ = db_limits.Incr(database.Ctx, "counter")

	http.Redirect(w, r, value, http.StatusSeeOther)
}
