package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Coins struct {
	ID                                string      `json:"id"`
	Symbol                            string      `json:"symbol"`
	Name                              string      `json:"name"`
	Image                             string      `json:"image"`
	CurrentPrice                      float64     `json:"current_price"`
	MarketCap                         int64       `json:"market_cap"`
	MarketCapRank                     float64     `json:"market_cap_rank"`
	FullyDilutedValuation             int64       `json:"fully_diluted_valuation"`
	TotalVolume                       int64       `json:"total_volume"`
	High24H                           float64     `json:"high_24h"`
	Low24H                            float64     `json:"low_24h"`
	PriceChange24H                    float64     `json:"price_change_24h"`
	PriceChangePercentage24H          float64     `json:"price_change_percentage_24h"`
	MarketCapChange24H                float64     `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H      float64     `json:"market_cap_change_percentage_24h"`
	CirculatingSupply                 float64     `json:"circulating_supply"`
	TotalSupply                       float64     `json:"total_supply"`
	MaxSupply                         float64     `json:"max_supply"`
	Ath                               float64     `json:"ath"`
	AthChangePercentage               float64     `json:"ath_change_percentage"`
	AthDate                           time.Time   `json:"ath_date"`
	Atl                               float64     `json:"atl"`
	AtlChangePercentage               float64     `json:"atl_change_percentage"`
	AtlDate                           time.Time   `json:"atl_date"`
	Roi                               interface{} `json:"roi"`
	LastUpdated                       time.Time   `json:"last_updated"`
	PriceChangePercentage1HInCurrency float64     `json:"price_change_percentage_1h_in_currency"`
}

func CoinsMarket() ([]Coins, error) {

	request, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=THB&order=market_cap_desc&per_page=10&page=1&sparkline=false&price_change_percentage=1h")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer request.Body.Close()

	var market []Coins

	if request.StatusCode == http.StatusOK {
		if err := json.Unmarshal(body, &market); err != nil {
			log.Fatalln(err)
			return market, err
		}
	}
	return market, nil
}
