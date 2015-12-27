package currency

import (
	"io/ioutil"
	"log"
	"mime"

	"github.com/hanksudo/bot-currency/lib"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Renew() {
	resp := lib.Fetch("http://rate.bot.com.tw/Pages/Static/UIP003.zh-TW.htm")
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	csvUri := lib.ExtractDownloadUrl(string(content))

	resp = lib.Fetch(csvUri)
	content_csv, err := ioutil.ReadAll(transform.NewReader(resp.Body, traditionalchinese.Big5.NewDecoder()))
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
			ioutil.WriteFile("csvs/"+params["filename"], content_csv, 0644)
			ioutil.WriteFile("latest.dat", []byte(params["filename"]), 0644)
		}
	}
}
