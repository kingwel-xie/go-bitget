package v2ext

import (
	"encoding/json"
	"github.com/kingwel-xie/go-bitget/internal/common"
	"github.com/kingwel-xie/go-bitget/internal/model"
	"github.com/kingwel-xie/go-bitget/pkg/client/ws"
)

type TickerEvent struct {
	InstId          string `json:"instId"`
	LastPr          string `json:"lastPr"`
	BidPr           string `json:"bidPr"`
	AskPr           string `json:"askPr"`
	BidSz           string `json:"bidSz"`
	AskSz           string `json:"askSz"`
	Open24H         string `json:"open24h"`
	High24H         string `json:"high24h"`
	Low24H          string `json:"low24h"`
	Change24H       string `json:"change24h"`
	FundingRate     string `json:"fundingRate"`
	NextFundingTime int64  `json:"nextFundingTime,string"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	HoldingAmount   string `json:"holdingAmount"`
	BaseVolume      string `json:"baseVolume"`
	QuoteVolume     string `json:"quoteVolume"`
	OpenUtc         string `json:"openUtc"`
	SymbolType      string `json:"symbolType"`
	Symbol          string `json:"symbol"`
	DeliveryPrice   string `json:"deliveryPrice"`
	Ts              int64  `json:"ts,string"`
}

// WsTickerHandler handle ticker event
type WsTickerHandler func([]TickerEvent)

// WsServeTickerStream retrieve the latest ticker data  of the instruments.
// instType
// SPOT 现货
// USDT-FUTURES USDT专业合约
// COIN-FUTURES 混合合约
// USDC-FUTURES USDC专业合约
// SUSDT-FUTURES USDT专业合约模拟盘
// SCOIN-FUTURES 混合合约模拟盘
// SUSDC-FUTURES USDC专业合约模拟盘
func WsServeTickerStream(instType string, symbols []string, handler WsTickerHandler, errHandler ErrHandler) (chan struct{}, chan struct{}, error) {
	wsHandler := func(message string) {
		var event struct {
			GenericMessage
			Data []TickerEvent `json:"data"`
		}
		err := json.Unmarshal([]byte(message), &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event.Data)
	}
	client, doneCh, ctrlCh, err := new(ws.BitgetWsClient).Init(false, wsHandler, common.OnError(errHandler))
	if err != nil {
		return nil, nil, err
	}

	var channelsDef []model.SubscribeReq
	for _, s := range symbols {
		req := model.SubscribeReq{
			InstType: instType,
			Channel:  "ticker",
			InstId:   s,
		}
		channelsDef = append(channelsDef, req)
	}
	client.SubscribeDef(channelsDef)
	return doneCh, ctrlCh, nil
}
