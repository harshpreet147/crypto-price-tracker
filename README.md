﻿# Crypto Price Tracker

﻿# Crypto Price Tracker

## Docker Image
```
docker run -p 7575:7575 jasbirnetwork/cryptopricetracker:v1
```

## Access Crypto Price Tracker App

Open [http://localhost:7575/price](http://localhost:7575/price) with your browser to see the result.

## You should get reponse like below:

```
{
  "rates": {
    "BTC": 51117.08414254509,
    "ETH": 2826.057988625188,
    "XRP": 0.8513239559964851
  }
}
```

## Explaination

#####  Api we have used
```
"https://rest.coinapi.io/v1/exchangerate/%s/CAD?apikey=%s"
```
* The first parameter we have sent coin type  `assets := []string{"BTC", "ETH", "XRP"}`.
* Second parameter it required is apikey `?apikey=%s`

```
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}
```
### Get and return the api response  
```
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
```
*  Prepare the string array to get type of asset `assets := []string{"BTC", "ETH", "XRP"}`
*  Perform the `GET` request inside the loop to get current price of crypto that we have defined in the prevoius step and append in the rate struct. 
* If reading rate return any error than it will return `fmt.Errorf`.

## Preparing response

```
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
```
* After call the function `fetchExchangeRates` and get response, set the header to JSON and send the response to user.

## Run `http` server and define api endpoint `/price`

```
func main() {
	http.HandleFunc("/price", getPriceHandler)

	port := ":7575" // Set the port
	log.Printf("Server listening on port %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
```
* Application will expose on port number `:7575`
* Open [http://localhost:7575/price](http://localhost:7575/price) with your browser to see the result.

### For test the app run below command.

```
go test
```
