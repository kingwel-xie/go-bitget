package v2ext

import (
	"fmt"
	"github.com/kingwel-xie/go-bitget/config"
	"os"
	"testing"
)

func init() {
	config.Init("bg_96f118260b9e5abda0a862fb0a6e6ffe", "00a850a6e0bd5bcfa61e5c0f86fe9fbe4213de458d13866d11046d47e3b3abd5", "pleasekeepitsafe")
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:8011")
	os.Setenv("HTTPs_PROXY", "http://127.0.0.1:8011")
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
