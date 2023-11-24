package main

import (
	"log"
	"net/http"
)
const coinAPIBaseURL = "https://rest.coinapi.io/v1/exchangerate/%s/CAD?apikey=%s"

type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

func fetchExchangeRates(apiKey string) (ExchangeRates, error) {
assets := []string{"BTC", "ETH", "XRP"}
	rates := make(map[string]float64)

	for _, asset := range assets {
		url := fmt.Sprintf(coinAPIBaseURL, asset, apiKey)
		resp, err := http.Get(url)
		if err != nil {
			return ExchangeRates{}, err
		}
		defer resp.Body.Close()

		var data map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return ExchangeRates{}, err
		}

		rate, ok := data["rate"].(float64)
		if !ok {
			return ExchangeRates{}, fmt.Errorf("unexpected response format for asset %s", asset)
		}

		rates[asset] = rate
	}

	return ExchangeRates{Rates: rates}, nil
}

func getPriceHandler(w http.ResponseWriter, r *http.Request) {

}

func getPriceHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := "FD1CD239-A4FC-4E9C-9485-3657E01DBDB9"
	exchangeRates, err := fetchExchangeRates(apiKey)
	if err != nil {
		http.Error(w, "Failed to fetch exchange rates", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exchangeRates)
}

func main() {
	http.HandleFunc("/price", getPriceHandler)

	port := ":7575" // Set the port
	log.Printf("Server listening on port %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
