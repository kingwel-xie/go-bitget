package v2ext

import (
	"fmt"
	"testing"
)

func TestAccountInfo(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.AccountInfo()
	fmt.Println(res)
}

func TestUserAssets(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.Assets("hold_only")
	fmt.Println(res)
}

func TestSpotCoins(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.Coins("")
	fmt.Println(res)
}

func TestSpotSymbols(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.Symbols("")
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

func TestSpotCandles(t *testing.T) {
	c := NewSpotClient()
	res, _ := c.Candles("DOTUSDT", "1m", 0, 0, 1000)
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

func TestSpotTickerStream(t *testing.T) {
	WsServeTickerStream([]string{"DOTUSDT", "ETHUSDT"}, func(events []SpotTickerEvent) {
		fmt.Println(events)
	}, func(err error) {
		fmt.Println(err)
	})

	select {}
}

func TestSpotUserDataStream(t *testing.T) {
	WsServeDataStream(func(event WsUserDataEvent) {
		fmt.Println(event)
	}, func(err error) {
		fmt.Println(err)
	})

	select {}
}
