package lib

import (
	"html"
	"log"
	"net/http"
	"regexp"
)

func Fetch(url string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Language", "en")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func ExtractDownloadUrl(content string) string {
	re := regexp.MustCompile("<a id=\"DownloadCsv\" class=\"buttonLink\" href=\"(.+)\">")
	return "http://rate.bot.com.tw" + html.UnescapeString(re.FindStringSubmatch(content)[1])
}
