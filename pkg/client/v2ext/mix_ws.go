package v2ext

import (
	"encoding/json"
	"fmt"
	"github.com/kingwel-xie/go-bitget/internal/common"
	"github.com/kingwel-xie/go-bitget/internal/model"
	"github.com/kingwel-xie/go-bitget/pkg/client/ws"
)

type MixAccountUpdateEvent struct {
	MarginCoin          string `json:"marginCoin"`
	Frozen              string `json:"frozen"`
	Available           string `json:"available"`
	MaxOpenPosAvailable string `json:"maxOpenPosAvailable"`
	MaxTransferOut      string `json:"maxTransferOut"`
	Equity              string `json:"equity"`
	UsdtEquity          string `json:"usdtEquity"`
}

type MixFillUpdateEvent struct {
	OrderId     string `json:"orderId"`
	TradeId     string `json:"tradeId"`
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	OrderType   string `json:"orderType"`
	PosMode     string `json:"posMode"`
	Price       string `json:"price"`
	BaseVolume  string `json:"baseVolume"`
	QuoteVolume string `json:"quoteVolume"`
	Profit      string `json:"profit"`
	TradeSide   string `json:"tradeSide"`
	TradeScope  string `json:"tradeScope"`
	FeeDetail   []struct {
		FeeCoin           string `json:"feeCoin"`
		Deduction         string `json:"deduction"`
		TotalDeductionFee string `json:"totalDeductionFee"`
		TotalFee          string `json:"totalFee"`
	} `json:"feeDetail"`
	CTime string `json:"cTime"`
	UTime string `json:"uTime"`
}

type MixOrderUpdateEvent struct {
	AccBaseVolume string `json:"accBaseVolume"`
	CTime         string `json:"cTime"`
	ClientOId     string `json:"clientOId"`
	FeeDetail     []struct {
		FeeCoin string `json:"feeCoin"`
		Fee     string `json:"fee"`
	} `json:"feeDetail"`
	FillFee          string `json:"fillFee"`
	FillFeeCoin      string `json:"fillFeeCoin"`
	FillNotionalUsd  string `json:"fillNotionalUsd"`
	FillPrice        string `json:"fillPrice"`
	BaseVolume       string `json:"baseVolume"`
	FillTime         string `json:"fillTime"`
	Force            string `json:"force"`
	InstId           string `json:"instId"`
	Leverage         string `json:"leverage"`
	MarginCoin       string `json:"marginCoin"`
	MarginMode       string `json:"marginMode"`
	NotionalUsd      string `json:"notionalUsd"`
	OrderId          string `json:"orderId"`
	OrderType        string `json:"orderType"`
	Pnl              string `json:"pnl"`
	PosMode          string `json:"posMode"`
	PosSide          string `json:"posSide"`
	Price            string `json:"price"`
	PriceAvg         string `json:"priceAvg"`
	ReduceOnly       string `json:"reduceOnly"`
	Side             string `json:"side"`
	Size             string `json:"size"`
	EnterPointSource string `json:"enterPointSource"`
	Status           string `json:"status"`
	TradeScope       string `json:"tradeScope"`
	TradeId          string `json:"tradeId"`
	TradeSide        string `json:"tradeSide"`
	UTime            int64  `json:"uTime,string"`
}

type WsMixUserDataEvent struct {
	Type          string `json:"type"`
	Event         string `json:"event"`
	Action        string `json:"action"`
	Time          int64  `json:"ts"`
	AccountUpdate []MixAccountUpdateEvent
	FillUpdate    []MixFillUpdateEvent
	OrderUpdate   []MixOrderUpdateEvent
}

// WsMixUserDataHandler handle user data event of MIX
type WsMixUserDataHandler func(WsMixUserDataEvent)

func WsServeMixDataStream(instType string, handler WsMixUserDataHandler, errHandler ErrHandler) (chan struct{}, chan struct{}, error) {
	wsHandler := func(message string) {
		fmt.Println(message)
	}

	client, doneCh, ctrlCh, err := new(ws.BitgetWsClient).Init(true, wsHandler, common.OnError(errHandler))
	if err != nil {
		return nil, nil, err
	}

	wsAccountHandler := func(message string) {
		var event struct {
			GenericMessage
			Data []MixAccountUpdateEvent `json:"data"`
		}
		err := json.Unmarshal([]byte(message), &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(WsMixUserDataEvent{
			Type:          event.Arg.InstType,
			Event:         event.Arg.Channel,
			Action:        event.Action,
			Time:          event.Ts,
			AccountUpdate: event.Data,
		})
	}
	req := model.SubscribeReq{
		InstType: instType,
		Channel:  "account",
		Coin:     "default",
	}
	client.SubscribeOne(req, wsAccountHandler)

	wsOrderHandler := func(message string) {
		var event struct {
			GenericMessage
			Data []MixOrderUpdateEvent `json:"data"`
		}
		err := json.Unmarshal([]byte(message), &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(WsMixUserDataEvent{
			Type:        event.Arg.InstType,
			Event:       event.Arg.Channel,
			Action:      event.Action,
			Time:        event.Ts,
			OrderUpdate: event.Data,
		})
	}
	req = model.SubscribeReq{
		InstType: instType,
		Channel:  "orders",
		InstId:   "default", // will get all symbols' events
	}
	client.SubscribeOne(req, wsOrderHandler)

	wsFillHandler := func(message string) {
		var event struct {
			GenericMessage
			Data []MixFillUpdateEvent `json:"data"`
		}
		err := json.Unmarshal([]byte(message), &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(WsMixUserDataEvent{
			Type:       event.Arg.InstType,
			Event:      event.Arg.Channel,
			Action:     event.Action,
			Time:       event.Ts,
			FillUpdate: event.Data,
		})
	}
	req = model.SubscribeReq{
		InstType: instType,
		Channel:  "fill",
		InstId:   "default",
	}
	client.SubscribeOne(req, wsFillHandler)

	return doneCh, ctrlCh, nil
}
