package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Currency struct {
	Last   float64 `json:"last"`
	Buy    float64 `json:"buy"`
	Sell   float64 `json:"sell"`
	Symbol string  `json:"symbol"`
}

func FetchCotationBTC() (Currency, error) {

	url := "https://blockchain.info/ticker"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return Currency{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Currency{}, err
	}

	var data map[string]Currency
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	usd, ok := data["USD"]
	if !ok {
		log.Fatalf("USD data not found")
	}

	return usd, nil
}

type AtomData struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

func FetchCotationATOM() (Currency, error) {

	url := "https://api.coincap.io/v2/assets/cosmos"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return Currency{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Currency{}, err
	}

	var response struct {
		Data AtomData `json:"data"`
	}

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	atomPrice, err := strconv.ParseFloat(response.Data.PriceUsd, 64)
	if err != nil {
		fmt.Println("Error converting Atom price to float64:", err)
		return Currency{}, err
	}

	var usd Currency
	usd.Last = atomPrice

	return usd, nil
}
