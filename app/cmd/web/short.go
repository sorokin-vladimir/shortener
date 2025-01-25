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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if result.Err != nil {
		handleError(result.Err, w)
		return
	}

	fmt.Fprintf(w, `
		<span class="short-arrow">-></span>
		<input type="text" id="short" value="%s" class="input-short" />
		<button type="button" class="copy-short" id="copy-short" aria-label="Copy short link">
			<svg class="icon-copy" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><g id="copy_bgCarrier"></g><g id="copy_tracerCarrier"></g><g id="copy_iconCarrier"> <path d="M19.5 16.5L19.5 4.5L18.75 3.75H9L8.25 4.5L8.25 7.5L5.25 7.5L4.5 8.25V20.25L5.25 21H15L15.75 20.25V17.25H18.75L19.5 16.5ZM15.75 15.75L15.75 8.25L15 7.5L9.75 7.5V5.25L18 5.25V15.75H15.75ZM6 9L14.25 9L14.25 19.5L6 19.5L6 9Z"></path> </g></svg>
		</button>
	`, result.Short)
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
