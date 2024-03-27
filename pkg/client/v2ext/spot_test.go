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

func TestInfo(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.Info()
	fmt.Println(res)
}

func TestUserAssets(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.Assets("hold_only")
	fmt.Println(res)
}

func TestTransfer(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.Transfer("spot", "coin_futures", "0.0001", "DOT")
	fmt.Println(res)
}

func TestSpotTicker(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.Tickers("DOTUSDT")
	fmt.Println(res)
}

//func TestPlaceOrder(t *testing.T) {
//	c := NewSpotClient()
//	res, _ := c.PlaceOrder("COIN-FUTURES", "DOTUSD", "sell", "open", "limit", "crossed", "DOT", "1", "")
//	fmt.Println(res)
//}

func TestSavingsAccount(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.SavingsAccount()
	fmt.Println(res)
}

func TestSavingsAssets(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.SavingsAssets("flexible")
	fmt.Println(res)
}
