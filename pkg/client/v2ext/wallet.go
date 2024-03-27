package v2ext

import (
	"encoding/json"
	"github.com/kingwel-xie/go-bitget/internal"
)

type TransferResponse struct {
	TransferId string `json:"transferId"`
	ClientOid  string `json:"clientOid"`
}

// Transfer
// spot 现货账户
// p2p P2P货账户
// coin_futures 币本位合约账户
// usdt_futures U本位合约账户
// usdc_futures USDC合约账户
// crossed_margin 全仓杠杆账户
// isolated_margin 逐仓杠杆账户
func (p *SpotClient) Transfer(fromType, toType, amount, coin string) (*TransferResponse, error) {
	params := map[string]string{
		"fromType": fromType,
		"toType":   toType,
		"amount":   amount,
		"coin":     coin,
	}

	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/transfer", postBody)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data *TransferResponse
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}
