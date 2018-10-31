package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volumn           int
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}

func main() {

	runtime.GOMAXPROCS(4)

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	numComplete := 0

	for _, symbol := range stockSymbols {
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)

	}

	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

	elapse := time.Since(start)

	fmt.Printf("Excution time: %s", elapse)
}
