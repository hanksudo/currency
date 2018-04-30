package currency

import (
	"io/ioutil"
	"log"
	"mime"

	"github.com/hanksudo/bot-currency/lib"
)

// Renew - currency data
func Renew() {
	log.Println("Currency data renew")
	resp := lib.Fetch("http://rate.bot.com.tw/xrt/flcsv/0/day")
	contentCsv, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// save csv
	contentDisposition := resp.Header["Content-Disposition"]
	if len(contentDisposition) > 0 {
		mediatype, params, err := mime.ParseMediaType(contentDisposition[0])
		if err != nil {
			log.Fatal(err)
		}
		if mediatype == "attachment" {
			ioutil.WriteFile("csvs/"+params["filename"], contentCsv, 0644)
			ioutil.WriteFile("latest.dat", []byte(params["filename"]), 0644)
		}
	}
}
