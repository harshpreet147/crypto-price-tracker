package main

import (
	"log"
	"net/http"
)

func fetchExchangeRates(apiKey string) (ExchangeRates, error) {

}

func getPriceHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/price", getPriceHandler)

	port := ":7575" // Set the port
	log.Printf("Server listening on port %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
