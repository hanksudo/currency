package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

type SlackAttachment struct {
	Text string `json:"text"`
}

type SlackPostMessage struct {
	ResponseType string            `json:"response_type"`
	Text         string            `json:"text"`
	Attachments  []SlackAttachment `json:"attachments"`
}

func slack_handler(w http.ResponseWriter, r *http.Request) {
	currency, err := info.Get(r.FormValue("text"))
	response := ""
	if err != nil {
		response = "找不到這個貨幣耶"
	} else {
		response = fmt.Sprintf("現金買入: %v, 現金賣出: %v\n即期買入: %v, 即期賣出: %v", currency.BuyCach, currency.SellCash, currency.BuySpot, currency.SellSpot)
	}
	postMessage := SlackPostMessage{}
	postMessage.ResponseType = "in_channel"
	postMessage.Text = "Currency rate for: " + strings.ToUpper(r.FormValue("text"))

	attachment := SlackAttachment{Text: response}
	postMessage.Attachments = []SlackAttachment{attachment}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postMessage)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/slack", slack_handler)
	http.HandleFunc("/renew", renew_handler)
	log.Fatal(http.ListenAndServe(":3030", Log(http.DefaultServeMux)))
}
