package v2ext

import (
	"encoding/json"
	"github.com/kingwel-xie/go-bitget/internal"
)

// PlaceOrder normal order
// side: buy/sell, tradeSide: open/close, orderType: limit/market, marginMode: isolated/crossed
// 双向持仓时，开多规则为：side=buy,tradeSide=open；开空规则为：side=sell,tradeSide=open；平多规则为：side=buy,tradeSide=close；平空规则为：side=sell,tradeSide=close
// force: ioc 无法立即成交的部分就撤销, fok 无法全部立即成交就撤销, gtc 普通订单, 订单会一直有效，直到被成交或者取消, post_only 只做maker, 订单类型为限价单(limit)时必填，若省略则默认为gtc
func (p *MixClient) PlaceOrder(productType, symbol, side, tradeSide, orderType string, postOnly bool, marginMode, marginCoin string, size, price string) (*OrderResponse, error) {
	params := map[string]string{
		"productType": productType,
		"symbol":      symbol,
		"side":        side,
		"tradeSide":   tradeSide,
		"orderType":   orderType,
		"marginMode":  marginMode,
		"marginCoin":  marginCoin,
		"size":        size,
	}
	if orderType == "limit" {
		if postOnly {
			params["force"] = "post_only"
		} else {
			params["force"] = "gtc"
		}
		params["price"] = price
	}

	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.DoPost("/api/v2/mix/order/place-order", postBody)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data *OrderResponse
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

func (p *MixClient) CancelOrder(productType, symbol, orderId string) (*OrderResponse, error) {
	params := map[string]string{
		"productType": productType,
		"symbol":      symbol,
		"orderId":     orderId,
	}
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.DoPost("/api/v2/mix/order/cancel-order", postBody)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data *OrderResponse
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

type MixPendingOrderEntry struct {
	Symbol                 string `json:"symbol"`
	Size                   string `json:"size"`
	OrderId                string `json:"orderId"`
	ClientOid              string `json:"clientOid"`
	BaseVolume             string `json:"baseVolume"`
	Fee                    string `json:"fee"`
	Price                  string `json:"price"`
	PriceAvg               string `json:"priceAvg"`
	Status                 string `json:"status"`
	Side                   string `json:"side"`
	Force                  string `json:"force"`
	TotalProfits           string `json:"totalProfits"`
	PosSide                string `json:"posSide"`
	MarginCoin             string `json:"marginCoin"`
	QuoteVolume            string `json:"quoteVolume"`
	Leverage               string `json:"leverage"`
	MarginMode             string `json:"marginMode"`
	EnterPointSource       string `json:"enterPointSource"`
	TradeSide              string `json:"tradeSide"`
	PosMode                string `json:"posMode"`
	OrderType              string `json:"orderType"`
	OrderSource            string `json:"orderSource"`
	CTime                  int64  `json:"cTime,string"`
	UTime                  int64  `json:"uTime,string"`
	PresetStopSurplusPrice string `json:"presetStopSurplusPrice"`
	PresetStopLossPrice    string `json:"presetStopLossPrice"`
}

func (p *MixClient) PendingOrders(productType, symbol, orderId string) ([]*MixPendingOrderEntry, error) {
	params := map[string]string{
		"productType": productType,
		"symbol":      symbol,
		"orderId":     orderId,
	}
	resp, err := p.DoGet("/api/v2/mix/order/orders-pending", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data struct {
			EntrustedList []*MixPendingOrderEntry `json:"entrustedList"`
			EndId         string                  `json:"endId"`
		}
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data.EntrustedList, err
}
