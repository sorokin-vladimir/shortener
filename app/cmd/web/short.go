package web

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/sorokin-vladimir/shortener/internal/shortener"
	"github.com/sorokin-vladimir/shortener/internal/utils"
)

type RequestPayload struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

func short(w http.ResponseWriter, r *http.Request) {
	errMethod := utils.CheckHttpMethod(r.Method, http.MethodPost)
	if errMethod != nil {
		http.Error(w, errMethod.Error(), http.StatusMethodNotAllowed)
		log.Println(errMethod.Error())
		return
	}

	clientIP := getClientIP(r)

	if r.Header.Get("HX-Request") != "true" {
		http.Error(w, utils.FORM_EXPECTED, http.StatusBadRequest)
		log.Println(utils.FORM_EXPECTED)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, utils.FAILED_PARSING_FORM, http.StatusBadRequest)
		log.Println(utils.FAILED_PARSING_FORM)
		return
	}

	url := r.FormValue("url")

	result := shortener.Shortener(
		url,
		clientIP,
		0,
		0,
		"",
	)

	if result.Err != nil {
		http.Error(w, result.Err.Error(), http.StatusInternalServerError)
		log.Println(result.Err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<div>
			<span class="short" id="short">%s</span>
			<input type="text" id="short-input" value="%s" style="position:absolute;left:-1000000px;opacity:0;" />
			<button type="button" class="copy-short" id="copy-short">Copy</button>
		</div>
	`, result.Short, result.Short)
}

func getClientIP(r *http.Request) string {
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		return fwd
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "anonymous"
	}
	return ip
}
