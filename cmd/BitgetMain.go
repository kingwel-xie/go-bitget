package main

import (
	"fmt"
	"github.com/kingwel-xie/go-bitget/internal/model"
	"github.com/kingwel-xie/go-bitget/pkg/client/ws"
)

func main() {
	client := new(ws.BitgetWsClient).Init(true, func(message string) {
		fmt.Println("default error:" + message)
	}, func(message string) {
		fmt.Println("default error:" + message)
	})

	var channelsDef []model.SubscribeReq
	subReqDef1 := model.SubscribeReq{
		InstType: "SPOT",
		Channel:  "ticker",
		InstId:   "DOTUSDT",
	}
	channelsDef = append(channelsDef, subReqDef1)
	client.SubscribeDef(channelsDef)

	var channels []model.SubscribeReq
	subReq1 := model.SubscribeReq{
		InstType: "SPOT",
		Channel:  "trade",
		InstId:   "DOTUSDT",
	}
	channels = append(channels, subReq1)
	client.Subscribe(channels, func(message string) {
		fmt.Println("appoint:" + message)
	})
	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()
}
