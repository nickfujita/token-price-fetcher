package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Ticker struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Symbol           string `json:"symbol,omitempty"`
	Rank             string `json:"rank,omitempty"`
	Currency         string `json:"currency"` //internally assign field
	PriceUsd         string `json:"price_usd,omitempty"`
	PriceBtc         string `json:"price_btc,omitempty"`
	Two4HVolumeUsd   string `json:"24h_volume_usd,omitempty"`
	MarketCapUsd     string `json:"market_cap_usd,omitempty"`
	AvailableSupply  string `json:"available_supply,omitempty"`
	TotalSupply      string `json:"total_supply,omitempty"`
	PercentChange1H  string `json:"percent_change_1h,omitempty"`
	PercentChange24H string `json:"percent_change_24h,omitempty"`
	PercentChange7D  string `json:"percent_change_7d,omitempty"`
	LastUpdated      string `json:"last_updated,omitempty"`

	PriceJpy       string `json:"price_jpy,omitempty"`
	Two4HVolumeJpy string `json:"24h_volume_jpy,omitempty"`
	MarketCapJpy   string `json:"market_cap_jpy,omitempty"`

	Priceeur       string `json:"price_eur,omitempty"`
	Two4HVolumeeur string `json:"24h_volume_eur,omitempty"`
	MarketCapeur   string `json:"market_cap_eur,omitempty"`

	Priceaud       string `json:"price_aud,omitempty"`
	Two4HVolumeaud string `json:"24h_volume_aud,omitempty"`
	MarketCapaud   string `json:"market_cap_aud,omitempty"`

	Pricebrl       string `json:"price_brl,omitempty"`
	Two4HVolumebrl string `json:"24h_volume_brl,omitempty"`
	MarketCapbrl   string `json:"market_cap_brl,omitempty"`

	Pricecad       string `json:"price_cad,omitempty"`
	Two4HVolumecad string `json:"24h_volume_cad,omitempty"`
	MarketCapcad   string `json:"market_cap_cad,omitempty"`

	Pricechf       string `json:"price_chf,omitempty"`
	Two4HVolumechf string `json:"24h_volume_chf,omitempty"`
	MarketCapchf   string `json:"market_cap_chf,omitempty"`

	Priceclp       string `json:"price_clp,omitempty"`
	Two4HVolumeclp string `json:"24h_volume_clp,omitempty"`
	MarketCapclp   string `json:"market_cap_clp,omitempty"`

	Pricecny       string `json:"price_cny,omitempty"`
	Two4HVolumecny string `json:"24h_volume_cny,omitempty"`
	MarketCapcny   string `json:"market_cap_cny,omitempty"`

	Priceczk       string `json:"price_czk,omitempty"`
	Two4HVolumeczk string `json:"24h_volume_czk,omitempty"`
	MarketCapczk   string `json:"market_cap_czk,omitempty"`

	Pricedkk       string `json:"price_dkk,omitempty"`
	Two4HVolumedkk string `json:"24h_volume_dkk,omitempty"`
	MarketCapdkk   string `json:"market_cap_dkk,omitempty"`

	Pricegbp       string `json:"price_gbp,omitempty"`
	Two4HVolumegbp string `json:"24h_volume_gbp,omitempty"`
	MarketCapgbp   string `json:"market_cap_gbp,omitempty"`

	Pricehkd       string `json:"price_hkd,omitempty"`
	Two4HVolumehkd string `json:"24h_volume_hkd,omitempty"`
	MarketCaphkd   string `json:"market_cap_hkd,omitempty"`

	Pricehuf       string `json:"price_huf,omitempty"`
	Two4HVolumehuf string `json:"24h_volume_huf,omitempty"`
	MarketCaphuf   string `json:"market_cap_huf,omitempty"`

	Priceidr       string `json:"price_idr,omitempty"`
	Two4HVolumeidr string `json:"24h_volume_idr,omitempty"`
	MarketCapidr   string `json:"market_cap_idr,omitempty"`

	Priceils       string `json:"price_ils,omitempty"`
	Two4HVolumeils string `json:"24h_volume_ils,omitempty"`
	MarketCapils   string `json:"market_cap_ils,omitempty"`

	Priceinr       string `json:"price_inr,omitempty"`
	Two4HVolumeinr string `json:"24h_volume_inr,omitempty"`
	MarketCapinr   string `json:"market_cap_inr,omitempty"`

	Pricekrw       string `json:"price_krw,omitempty"`
	Two4HVolumekrw string `json:"24h_volume_krw,omitempty"`
	MarketCapkrw   string `json:"market_cap_krw,omitempty"`

	Pricemxn       string `json:"price_mxn,omitempty"`
	Two4HVolumemxn string `json:"24h_volume_mxn,omitempty"`
	MarketCapmxn   string `json:"market_cap_mxn,omitempty"`

	Pricemyr       string `json:"price_myr,omitempty"`
	Two4HVolumemyr string `json:"24h_volume_myr,omitempty"`
	MarketCapmyr   string `json:"market_cap_myr,omitempty"`

	Pricenok       string `json:"price_nok,omitempty"`
	Two4HVolumenok string `json:"24h_volume_nok,omitempty"`
	MarketCapnok   string `json:"market_cap_nok,omitempty"`

	Pricenzd       string `json:"price_nzd,omitempty"`
	Two4HVolumenzd string `json:"24h_volume_nzd,omitempty"`
	MarketCapnzd   string `json:"market_cap_nzd,omitempty"`

	Pricephp       string `json:"price_php,omitempty"`
	Two4HVolumephp string `json:"24h_volume_php,omitempty"`
	MarketCapphp   string `json:"market_cap_php,omitempty"`

	Pricepkr       string `json:"price_pkr,omitempty"`
	Two4HVolumepkr string `json:"24h_volume_pkr,omitempty"`
	MarketCappkr   string `json:"market_cap_pkr,omitempty"`

	Pricepln       string `json:"price_pln,omitempty"`
	Two4HVolumepln string `json:"24h_volume_pln,omitempty"`
	MarketCappln   string `json:"market_cap_pln,omitempty"`

	Pricerub       string `json:"price_rub,omitempty"`
	Two4HVolumerub string `json:"24h_volume_rub,omitempty"`
	MarketCaprub   string `json:"market_cap_rub,omitempty"`

	Pricesek       string `json:"price_sek,omitempty"`
	Two4HVolumesek string `json:"24h_volume_sek,omitempty"`
	MarketCapsek   string `json:"market_cap_sek,omitempty"`

	Pricesgd       string `json:"price_sgd,omitempty"`
	Two4HVolumesgd string `json:"24h_volume_sgd,omitempty"`
	MarketCapsgd   string `json:"market_cap_sgd,omitempty"`

	Pricethb       string `json:"price_thb,omitempty"`
	Two4HVolumethb string `json:"24h_volume_thb,omitempty"`
	MarketCapthb   string `json:"market_cap_thb,omitempty"`

	Pricetry       string `json:"price_try,omitempty"`
	Two4HVolumetry string `json:"24h_volume_try,omitempty"`
	MarketCaptry   string `json:"market_cap_try,omitempty"`

	Pricetwd       string `json:"price_twd,omitempty"`
	Two4HVolumetwd string `json:"24h_volume_twd,omitempty"`
	MarketCaptwd   string `json:"market_cap_twd,omitempty"`

	Pricezar       string `json:"price_zar,omitempty"`
	Two4HVolumezar string `json:"24h_volume_zar,omitempty"`
	MarketCapzar   string `json:"market_cap_zar,omitempty"`
}

func (t *Ticker) Bytes() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		return nil
	}
	return b
}

func doEvery(d time.Duration, f func(string, string, time.Time), symbols []string, currencies []string) {

	for t := range time.Tick(d) {
		for _, s := range symbols {
			for _, c := range currencies {
				f(s, c, t)
			}
		}
	}
}

func fetch(symbol string, currency string, t time.Time) {
	url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s/?convert=%s", symbol, currency)
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	tickers := []Ticker{}

	err = json.NewDecoder(res.Body).Decode(&tickers)
	if err != nil {
		log.Printf("error decode %v", err)
		return
	}

	for _, v := range tickers {
		v.Currency = currency
		log.Printf("%v\n", v)
		//save your ticker result here
	}
}

func main() {

	currencies := []string{"USD", "AUD", "BRL", "CAD", "CHF", "CLP", "CNY", "CZK", "DKK", "EUR", "GBP", "HKD", "HUF", "IDR", "ILS", "INR", "JPY", "KRW", "MXN", "MYR", "NOK", "NZD", "PHP", "PKR", "PLN", "RUB", "SEK", "SGD", "THB", "TRY", "TWD", "ZAR"}
	symbols := []string{"neo", "gas", "red-pulse"}
	for _, s := range symbols {
		for _, c := range currencies {
			fetch(s, c, time.Now())
		}
	}

	doEvery(2*time.Minute, fetch, symbols, currencies)
}
