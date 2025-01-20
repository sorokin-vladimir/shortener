package web

import (
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/sorokin-vladimir/shortener/internal/database"
	"github.com/sorokin-vladimir/shortener/internal/utils"
)

func resolveShort(w http.ResponseWriter, r *http.Request) {
	errMethod := utils.CheckHttpMethod(r.Method, http.MethodGet)
	if errMethod != nil {
		http.Error(w, errMethod.Error(), http.StatusMethodNotAllowed)
		log.Println(errMethod.Error())
		return
	}

	db_shorts := database.CreateClient(database.DB_SHORT)
	defer db_shorts.Close()

	short := r.PathValue("short")
	value, err := db_shorts.Get(database.Ctx, short).Result()

	if err == redis.Nil {
		log.Printf("Could not resolve '%s'", short)
		http.Error(w, "Could not resolve "+short, http.StatusNotFound)
		return
	} else if err != nil {
		log.Println("Internal error", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	db_limits := database.CreateClient(database.DB_LIMITS)
	defer db_limits.Close()

	_ = db_limits.Incr(database.Ctx, "counter")

	log.Printf("Resolved '%s' to '%s'", short, value)
	http.Redirect(w, r, value, http.StatusSeeOther)
}
