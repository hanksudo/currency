package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/hanksudo/bot-currency/currency"
	"github.com/hanksudo/bot-currency/info"
)

func logging(handler http.Handler) http.Handler {
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

func renewHandler(w http.ResponseWriter, r *http.Request) {
	currency.Renew()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

type slackAttachment struct {
	Text string `json:"text"`
}

type slackPostMessage struct {
	ResponseType string            `json:"response_type"`
	Text         string            `json:"text"`
	Attachments  []slackAttachment `json:"attachments"`
}

func slackHandler(w http.ResponseWriter, r *http.Request) {
	currency, err := info.Get(r.FormValue("text"))
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

func main() {
	webPtr := flag.Bool("web", false, "Start web server")
	renewPtr := flag.Bool("renew", false, "Renew currency data")
	backupPtr := flag.Bool("backup", false, "Backup to Dropbox")
	flag.Parse()

	if *webPtr {
		startWebService()
	} else if *renewPtr {
		currency.Renew()
	} else if *backupPtr {
		fmt.Println("Start backup")
		cmd := exec.Command("python", "scripts/backup_to_dropbox.py")
		outPipe, _ := cmd.StdoutPipe()
		errPipe, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(outPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		scanner = bufio.NewScanner(errPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		cmd.Wait()
	}
}

func startWebService() {
	log.Println("Start web service.")
	http.HandleFunc("/", handler)
	http.HandleFunc("/slack", slackHandler)
	http.HandleFunc("/renew", renewHandler)
	log.Fatal(http.ListenAndServe(":3030", logging(http.DefaultServeMux)))
}
