package v2ext

import (
	"encoding/json"
	"github.com/kingwel-xie/go-bitget/internal"
)

// PlaceOrder normal order
// side: buy/sell, tradeSide: open/close, orderType: limit/market, marginMode: isolated/crossed
// 双向持仓时，开多规则为：side=buy,tradeSide=open；开空规则为：side=sell,tradeSide=open；平多规则为：side=buy,tradeSide=close；平空规则为：side=sell,tradeSide=close
func (p *SpotClient) PlaceOrder(symbol, side, orderType string, postOnly bool, size, price string) (*OrderResponse, error) {
	params := map[string]string{
		"symbol":    symbol,
		"side":      side,
		"orderType": orderType,
		"size":      size,
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
	resp, err := p.DoPost("/api/v2/spot/trade/place-order", postBody)
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

func (p *SpotClient) CancelOrder(symbol, orderId string) (*OrderResponse, error) {
	params := map[string]string{
		"symbol":  symbol,
		"orderId": orderId,
	}
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.DoPost("/api/v2/spot/trade/cancel-order", postBody)
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

type SpotPendingOrderEntry struct {
	UserId           string      `json:"userId"`
	Symbol           string      `json:"symbol"`
	OrderId          string      `json:"orderId"`
	ClientOid        string      `json:"clientOid"`
	PriceAvg         string      `json:"priceAvg"`
	Size             string      `json:"size"`
	OrderType        string      `json:"orderType"`
	Side             string      `json:"side"`
	Status           string      `json:"status"`
	BasePrice        string      `json:"basePrice"`
	BaseVolume       string      `json:"baseVolume"`
	QuoteVolume      string      `json:"quoteVolume"`
	EnterPointSource string      `json:"enterPointSource"`
	TriggerPrice     interface{} `json:"triggerPrice"`
	TpslType         string      `json:"tpslType"`
	CTime            int64       `json:"cTime,string"`
}

func (p *SpotClient) PendingOrders(symbol, orderId string) ([]*SpotPendingOrderEntry, error) {
	params := map[string]string{
		"symbol":  symbol,
		"orderId": orderId,
	}
	resp, err := p.DoGet("/api/v2/spot/trade/unfilled-orders", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*SpotPendingOrderEntry
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}
