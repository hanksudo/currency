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

func slack_handler(w http.ResponseWriter, r *http.Request) {
	currency, err := info.Get(r.FormValue("text"))
	response := ""
	if err != nil {
		response = "找不到這個貨幣耶"
	} else {
		response = "現金買入: "+currency.BuyCach+", 現金賣出: "+currency.SellCash+"\n即期買入: "+currency.BuySpot+", 即期賣出: "+currency.SellSpot
	}
	json.NewEncoder(w).Encode(map[string]string{
	    "response_type": "in_channel",
	    "text": "Currency rate for: " strings.ToUpper(r.FormValue("text")),
	    "attachments": [
	        {
	            "text": response
	        }
	    ]
	})
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/slack", slack_handler)
	http.HandleFunc("/renew", renew_handler)
	log.Fatal(http.ListenAndServe(":3030", Log(http.DefaultServeMux)))
}
