package v2ext

import (
	"encoding/json"
	"fmt"
	"github.com/kingwel-xie/go-bitget/internal/common"
	"github.com/kingwel-xie/go-bitget/internal/model"
	"github.com/kingwel-xie/go-bitget/pkg/client/ws"
)

//
//type SpotTickerEvent struct {
//	InstID       string `json:"instId"`
//	LastPr       string `json:"lastPr"`
//	Open24H      string `json:"open24h"`
//	High24H      string `json:"high24h"`
//	Low24H       string `json:"low24h"`
//	Change24H    string `json:"change24h"`
//	BidPr        string `json:"bidPr"`
//	AskPr        string `json:"askPr"`
//	BidSz        string `json:"bidSz"`
//	AskSz        string `json:"askSz"`
//	BaseVolume   string `json:"baseVolume"`
//	QuoteVolume  string `json:"quoteVolume"`
//	OpenUtc      string `json:"openUtc"`
//	ChangeUtc24H string `json:"changeUtc24h"`
//	Ts           int64  `json:"ts,string"`
//}
//
//// WsTickerHandler handle ticker event
//type WsTickerHandler func([]SpotTickerEvent)
//
//func WsServeTickerStream(symbols []string, handler WsTickerHandler, errHandler ErrHandler) (chan struct{}, chan struct{}, error) {
//	wsHandler := func(message string) {
//		var event struct {
//			GenericMessage
//			Data []SpotTickerEvent `json:"data"`
//		}
//		err := json.Unmarshal([]byte(message), &event)
//		if err != nil {
//			errHandler(err)
//			return
//		}
//		handler(event.Data)
//	}
//	client, doneCh, ctrlCh, err := new(ws.BitgetWsClient).Init(false, wsHandler, common.OnError(errHandler))
//	if err != nil {
//		return nil, nil, err
//	}
//
//	var channelsDef []model.SubscribeReq
//	for _, s := range symbols {
//		req := model.SubscribeReq{
//			InstType: "SPOT",
//			Channel:  "ticker",
//			InstId:   s,
//		}
//		channelsDef = append(channelsDef, req)
//	}
//	client.SubscribeDef(channelsDef)
//	return doneCh, ctrlCh, nil
//}

type AccountUpdateEvent struct {
	Coin           string `json:"coin"`
	Available      string `json:"available"`
	Frozen         string `json:"frozen"`
	Locked         string `json:"locked"`
	LimitAvailable string `json:"limitAvailable"`
	UTime          int64  `json:"uTime,string"`
}

type FillUpdateEvent struct {
	OrderID    string `json:"orderId"`
	TradeID    string `json:"tradeId"`
	Symbol     string `json:"symbol"`
	OrderType  string `json:"orderType"`
	Side       string `json:"side"`
	PriceAvg   string `json:"priceAvg"`
	Size       string `json:"size"`
	Amount     string `json:"amount"`
	TradeScope string `json:"tradeScope"`
	FeeDetail  []struct {
		FeeCoin           string `json:"feeCoin"`
		Deduction         string `json:"deduction"`
		TotalDeductionFee string `json:"totalDeductionFee"`
		TotalFee          string `json:"totalFee"`
	} `json:"feeDetail"`
	CTime int64 `json:"cTime,string"`
	UTime int64 `json:"uTime,string"`
}

type OrderUpdateEvent struct {
	InstID        string `json:"instId"`
	OrderID       string `json:"orderId"`
	ClientOid     string `json:"clientOid"`
	Size          string `json:"size"`
	NewSize       string `json:"newSize"`
	Notional      string `json:"notional"`
	OrderType     string `json:"orderType"`
	Force         string `json:"force"`
	Side          string `json:"side"`
	FillPrice     string `json:"fillPrice"`
	TradeID       string `json:"tradeId"`
	BaseVolume    string `json:"baseVolume"`
	FillTime      int64  `json:"fillTime,string"`
	FillFee       string `json:"fillFee"`
	FillFeeCoin   string `json:"fillFeeCoin"`
	TradeScope    string `json:"tradeScope"`
	AccBaseVolume string `json:"accBaseVolume"`
	PriceAvg      string `json:"priceAvg"`
	Status        string `json:"status"`
	CTime         int64  `json:"cTime,string"`
	UTime         int64  `json:"uTime,string"`
	FeeDetail     []struct {
		FeeCoin string `json:"feeCoin"`
		Fee     string `json:"fee"`
	} `json:"feeDetail"`
	EnterPointSource string `json:"enterPointSource"`
}

type WsUserDataEvent struct {
	Type          string `json:"type"`
	Event         string `json:"event"`
	Action        string `json:"action"`
	Time          int64  `json:"ts"`
	AccountUpdate []AccountUpdateEvent
	FillUpdate    []FillUpdateEvent
	OrderUpdate   []OrderUpdateEvent
}

// WsUserDataHandler handle user data event
type WsUserDataHandler func(WsUserDataEvent)

func WsServeDataStream(handler WsUserDataHandler, errHandler ErrHandler) (chan struct{}, chan struct{}, error) {
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
			Data []AccountUpdateEvent `json:"data"`
		}
		err := json.Unmarshal([]byte(message), &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(WsUserDataEvent{
			Type:          event.Arg.InstType,
			Event:         event.Arg.Channel,
			Action:        event.Action,
			Time:          event.Ts,
			AccountUpdate: event.Data,
		})
	}
	req := model.SubscribeReq{
		InstType: "SPOT",
		Channel:  "account",
		Coin:     "default",
	}
	client.SubscribeOne(req, wsAccountHandler)

	wsOrderHandler := func(message string) {
		var event struct {
			GenericMessage
			Data []OrderUpdateEvent `json:"data"`
		}
		err := json.Unmarshal([]byte(message), &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(WsUserDataEvent{
			Type:        event.Arg.InstType,
			Event:       event.Arg.Channel,
			Action:      event.Action,
			Time:        event.Ts,
			OrderUpdate: event.Data,
		})
	}
	req = model.SubscribeReq{
		InstType: "SPOT",
		Channel:  "orders",
		InstId:   "default", // will get all symbols' events
	}
	client.SubscribeOne(req, wsOrderHandler)

	wsFillHandler := func(message string) {
		var event struct {
			GenericMessage
			Data []FillUpdateEvent `json:"data"`
		}
		err := json.Unmarshal([]byte(message), &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(WsUserDataEvent{
			Type:       event.Arg.InstType,
			Event:      event.Arg.Channel,
			Action:     event.Action,
			Time:       event.Ts,
			FillUpdate: event.Data,
		})
	}
	req = model.SubscribeReq{
		InstType: "SPOT",
		Channel:  "fill",
		InstId:   "default",
	}
	client.SubscribeOne(req, wsFillHandler)

	return doneCh, ctrlCh, nil
}
