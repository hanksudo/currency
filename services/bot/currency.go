package bot

import (
	"io"
	"log"
	"mime"
	"net/http"
	"os"
)

// Renew - currency data
func Renew() {
	log.Println("Refetch Currency data")
	resp := fetch("https://rate.bot.com.tw/xrt/flcsv/0/day")
	contentCsv, err := io.ReadAll(resp.Body)
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
		filename := params["filename"]
		if mediatype == "attachment" {
			if pathExists("csvs/" + filename) {
				log.Println("File exists", filename)
				return
			}
			os.WriteFile("csvs/"+filename, contentCsv, 0644)
			os.WriteFile("latest.dat", []byte(filename), 0644)
		}
	}
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func fetch(url string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Language", "en")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	return resp
}
