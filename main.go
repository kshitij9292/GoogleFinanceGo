package GoogleFinanceGo

import (
	_ "fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func Get_price(Ticker string, Exchange string) float64 {
	Exchange = ":" + Exchange
	var static = "https://www.google.com/finance/quote/"
	var class = "YMlKec fxKbKc"
	var url = static + Ticker + Exchange
	res, err := http.Get(url)
	if err != nil {
		log.Println("invalid ticker")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading body")
	}
	var Doc = string(body)
	DocIndex := strings.Index(Doc, class)
	var NewStr strings.Builder
	for i := DocIndex + 18; i < len(Doc); i++ {
		if Doc[i] == '<' {
			break
		}
		NewStr.WriteByte(Doc[i])
	}
	var PriceString = NewStr.String()
	PriceString = strings.Replace(PriceString, ",", "", -1)
	PriceFloat, err := strconv.ParseFloat(PriceString, 32)
	if err != nil {
		log.Println("error parsing Price:" + Ticker)
	}
	roundedPrice := roundFloat(PriceFloat, 2)
	return roundedPrice
}
func Previous_close(Ticker string, Exchange string) float64 {
	Exchange = ":" + Exchange
	var static = "https://www.google.com/finance/quote/"
	var class = "class=\"P6K39c\""
	var url = static + Ticker + Exchange
	res, err := http.Get(url)
	if err != nil {
		log.Println("invalid ticker")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading body")
	}
	var Doc = string(body)
	DocIndex := strings.Index(Doc, class)
	var NewStr strings.Builder
	for i := DocIndex + 18; i < len(Doc); i++ {
		if Doc[i] == '<' {
			break
		}
		NewStr.WriteByte(Doc[i])
	}
	var PriceString = NewStr.String()
	PriceString = strings.Replace(PriceString, ",", "", -1)
	PriceFloat, err := strconv.ParseFloat(PriceString, 32)
	if err != nil {
		log.Println("error parsing Price:" + Ticker)
	}
	roundedPrice := roundFloat(PriceFloat, 2)
	return roundedPrice
}
func About(Ticker string, Exchange string) string {
	Exchange = ":" + Exchange
	var static = "https://www.google.com/finance/quote/"
	var class = "class=\"bLLb2d\""
	var url = static + Ticker + Exchange
	res, err := http.Get(url)
	if err != nil {
		log.Println("invalid ticker")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading body")
	}
	var Doc = string(body)
	DocIndex := strings.Index(Doc, class)
	var NewStr strings.Builder
	var sentence int = 0
	for i := DocIndex + 15; i < len(Doc); i++ {
		if Doc[i] == '.' || Doc[i] == ',' {
			sentence++
		}
		if sentence == 2 {
			NewStr.WriteByte(Doc[i])
			break
		}
		NewStr.WriteByte(Doc[i])
	}
	var about = NewStr.String()
	for i := 0; i < len(about); i++ {
		if about[i] == '<' {
			return "_"
		}
	}
	return about
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
