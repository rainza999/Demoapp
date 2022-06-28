package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Ticker struct {
	Symbol        string  `json:"symbol"`
	BaseVolume    float64 `json:"baseVolume"`
	Change        float64 `json:"change"`
	High24Hr      float64 `json:"high24hr"`
	HighestBid    float64 `json:"highestBid"`
	ID            float64 `json:"id"`
	IsFrozen      float64 `json:"isFrozen"`
	Last          float64 `json:"last"`
	Low24Hr       float64 `json:"low24hr"`
	LowestAsk     float64 `json:"lowestAsk"`
	PercentChange float64 `json:"percentChange"`
	PrevClose     float64 `json:"prevClose"`
	PrevOpen      float64 `json:"prevOpen"`
	QuoteVolume   float64 `json:"quoteVolume"`
}

var BaseURL = "https://api.bitkub.com"

func TickerBitkub(symbol string) ([]Ticker, error) {
	var (
		result map[string]interface{}
		data   []Ticker
	)

	endpoint := fmt.Sprintf("%s%s", BaseURL, "/api/market/ticker")
	if symbol != "" {
		endpoint = fmt.Sprintf("%s%s?sym=%s", BaseURL, "/api/market/ticker", symbol)
	}

	request, err := http.Get(endpoint)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer request.Body.Close()

	if request.StatusCode == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalln(err)
			return nil, err
		}
	}

	count := 0

	for i, v := range result {
		if count < 10 {
			data = append(data, Ticker{
				Symbol:        i,
				BaseVolume:    v.(map[string]interface{})["baseVolume"].(float64),
				Change:        v.(map[string]interface{})["change"].(float64),
				High24Hr:      v.(map[string]interface{})["high24hr"].(float64),
				HighestBid:    v.(map[string]interface{})["highestBid"].(float64),
				ID:            v.(map[string]interface{})["id"].(float64),
				IsFrozen:      v.(map[string]interface{})["isFrozen"].(float64),
				Last:          v.(map[string]interface{})["last"].(float64),
				Low24Hr:       v.(map[string]interface{})["low24hr"].(float64),
				LowestAsk:     v.(map[string]interface{})["lowestAsk"].(float64),
				PercentChange: v.(map[string]interface{})["percentChange"].(float64),
				PrevClose:     v.(map[string]interface{})["prevClose"].(float64),
				PrevOpen:      v.(map[string]interface{})["prevOpen"].(float64),
				QuoteVolume:   v.(map[string]interface{})["quoteVolume"].(float64),
			})
			count++
		}
	}
	return data, nil
}
