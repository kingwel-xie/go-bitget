package v2ext

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:9999")
	os.Setenv("HTTPs_PROXY", "http://127.0.0.1:9999")
}

func TestAccountFunc(t *testing.T) {
	c := NewMixClient()
	res, _ := c.Accounts("COIN-FUTURES")
	fmt.Println(res)
}

func TestFundingRate(t *testing.T) {
	c := NewMixClient()
	res, _ := c.HistoryFundingRate("COIN-FUTURES", "DOTUSD", 1, 100)
	fmt.Println(res)
}

func TestTicker(t *testing.T) {
	c := NewMixClient()
	res, _ := c.Tickers("COIN-FUTURES")
	fmt.Println(res)
}

func TestPlaceOrder(t *testing.T) {
	c := NewMixClient()
	res, _ := c.PlaceOrder("COIN-FUTURES", "DOTUSD", "sell", "open", "limit", "crossed", "DOT", "1", "")
	fmt.Println(res)
}
