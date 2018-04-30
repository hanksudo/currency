package currency

import (
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Currency - data structure
type Currency struct {
	Name     string  `json:"name"`
	BuyCach  float64 `json:"buy_cash"`
	BuySpot  float64 `json:"buy_spot"`
	SellCash float64 `json:"sell_cash"`
	SellSpot float64 `json:"sell_spot"`
}

// Currencies - array
type Currencies []Currency

// Get - get currency info
func Get(currencyName string) (*Currency, error) {
	content, err := ioutil.ReadFile("latest.dat")
	if err != nil {
		return nil, err
	}

	latestCsv := string(content)
	contentCsv, err := ioutil.ReadFile("csvs/" + latestCsv)
	if err != nil {
		return nil, errors.New("fail")
	}

	r := csv.NewReader(strings.NewReader(string(contentCsv)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
				record = record[:len(record)-1]
			} else {
				log.Fatal(err)
			}
		}

		if strings.TrimSpace(record[0]) == strings.ToUpper(currencyName) {
			buyCash, _ := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)
			buySpot, _ := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
			sellCash, _ := strconv.ParseFloat(strings.TrimSpace(record[12]), 64)
			sellSpot, _ := strconv.ParseFloat(strings.TrimSpace(record[13]), 64)
			currency := Currency{
				record[0],
				buyCash,
				buySpot,
				sellCash,
				sellSpot,
			}
			return &currency, nil
		}
	}
	return nil, errors.New("fail")
}
