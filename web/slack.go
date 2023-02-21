package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hanksudo/currency/services/bot"
)

type slackAttachment struct {
	Text string `json:"text"`
}

type slackPostMessage struct {
	ResponseType string            `json:"response_type"`
	Text         string            `json:"text"`
	Attachments  []slackAttachment `json:"attachments"`
}

func slackHandler(w http.ResponseWriter, r *http.Request) {
	currency, err := bot.Get(r.FormValue("text"))
	response := ""
	if err != nil {
		response = "找不到這個貨幣耶\nUSD, HKD, GBP, AUD, CAD, SGD, CHF, JPY, ZAR, SEK, NZD, THB, PHP, IDR, EUR, KRW, VND, MYR, CNY"
	} else {
		response = fmt.Sprintf("現金買入: %v, 現金賣出: %v\n即期買入: %v, 即期賣出: %v", currency.BuyCach, currency.SellCash, currency.BuySpot, currency.SellSpot)
	}
	postMessage := slackPostMessage{}
	postMessage.ResponseType = "in_channel"
	postMessage.Text = "Currency rate for: " + strings.ToUpper(r.FormValue("text"))

	attachment := slackAttachment{Text: response}
	postMessage.Attachments = []slackAttachment{attachment}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postMessage)
}
