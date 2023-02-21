package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hanksudo/currency/services/bot"
)

// Start - web service
func Start() {
	port := 3030
	addr := fmt.Sprintf(":%d", port)
	log.Println("Start web service on http://localhost" + addr)
	http.HandleFunc("/", handler)
	http.HandleFunc("/slack", slackHandler)
	log.Fatal(http.ListenAndServe(addr, logging(http.DefaultServeMux)))
}

func logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	currencyName := r.URL.Query().Get("currency")
	if currencyName == "" {
		fmt.Fprint(w, "Hello?")
		return
	}

	content, err := bot.Get(currencyName)
	if err != nil {
		fmt.Fprint(w, "What?")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(content)
}
