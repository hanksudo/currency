package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hanksudo/bot-currency/currency"
	"github.com/hanksudo/bot-currency/info"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	currencyName := ""
	for k := range r.URL.Query() {
		currencyName = k
		break
	}

	if currencyName == "" {
		fmt.Fprint(w, "Hello?")
		return
	}
	content, err := info.Get(currencyName)
	if err != nil {
		fmt.Fprint(w, "What?")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(content)
}

func renew_handler(w http.ResponseWriter, r *http.Request) {
	currency.Renew()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/renew", renew_handler)
	log.Fatal(http.ListenAndServe("localhost:3001", Log(http.DefaultServeMux)))
}
