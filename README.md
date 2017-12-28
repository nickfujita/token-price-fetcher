# token-price-fetcher
A simple Go code to fetch ticker detail from coinmarketcap every N minutes.

For the simplicity sake, everything is in one file.

Update the list of tokens you want to fetch in `symbols`.  
See the list of supported tokens [here](https://api.coinmarketcap.com/v1/ticker/).  

##### Sample
```json
 {
        "id": "ethereum", 
        "name": "Ethereum", 
        "symbol": "ETH", 
        "rank": "2", 
        "price_usd": "754.059", 
        "price_btc": "0.0493773", 
        "24h_volume_usd": "2060270000.0", 
        "market_cap_usd": "72850263054.0", 
        "available_supply": "96610826.0", 
        "total_supply": "96610826.0", 
        "max_supply": null, 
        "percent_change_1h": "-1.14", 
        "percent_change_24h": "-3.04", 
        "percent_change_7d": "-9.19", 
        "last_updated": "1514422148"
    }, 
```

Note: use `id` field for symbol name

```go
func main() {
	currencies := []string{"USD", "AUD", "BRL", "CAD", "CHF", "CLP", "CNY", "CZK", "DKK", "EUR", "GBP", "HKD", "HUF", "IDR", "ILS", "INR", "JPY", "KRW", "MXN", "MYR", "NOK", "NZD", "PHP", "PKR", "PLN", "RUB", "SEK", "SGD", "THB", "TRY", "TWD", "ZAR"}
	symbols := []string{"neo", "gas", "red-pulse"}
	for _, s := range symbols {
		for _, c := range currencies {
			fetch(s, c, time.Now())
		}
	}
        //fetch every 5 minutes. Change to interval you want.
	doEvery(5*time.Minute, fetch, symbols, currencies) 
}

```

To run  
> go run main.go

