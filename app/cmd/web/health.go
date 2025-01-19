package web

import (
	"log"
	"net/http"

	"github.com/sorokin-vladimir/shortener/internal/database"
	"github.com/sorokin-vladimir/shortener/internal/utils"
)

func health(w http.ResponseWriter, r *http.Request) {
	errMethod := utils.CheckHttpMethod(r.Method, http.MethodGet)
	if errMethod != nil {
		http.Error(w, errMethod.Error(), http.StatusMethodNotAllowed)
		log.Println(errMethod.Error())
		return
	}

	db_shorts := database.CreateClient(database.DB_SHORT)
	defer db_shorts.Close()

	keys, err := db_shorts.Keys(database.Ctx, "*").Result()

	if err != nil {
		log.Println("Health check: Internal error", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	if len(keys) >= 0 {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
		return
	}
}
