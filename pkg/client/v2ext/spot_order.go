package v2ext

import (
	"encoding/json"
	"github.com/kingwel-xie/go-bitget/internal"
)

// PlaceOrder normal order
// side: buy/sell, tradeSide: open/close, orderType: limit/market, marginMode: isolated/crossed
// 双向持仓时，开多规则为：side=buy,tradeSide=open；开空规则为：side=sell,tradeSide=open；平多规则为：side=buy,tradeSide=close；平空规则为：side=sell,tradeSide=close
func (p *SpotClient) PlaceOrder(symbol, side, orderType string, force string, size, price string) (*OrderResponse, error) {
	params := map[string]string{
		"symbol":    symbol,
		"side":      side,
		"orderType": orderType,
		"size":      size,
	}
	if orderType == "limit" {
		params["force"] = force
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
