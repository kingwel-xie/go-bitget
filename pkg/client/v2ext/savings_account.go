package v2ext

import (
	"encoding/json"
	"github.com/kingwel-xie/go-bitget/internal"
	"strconv"
)

type SavingsAccountInfo struct {
	BtcAmount        string `json:"btcAmount"`
	UsdtAmount       string `json:"usdtAmount"`
	Btc24HEarning    string `json:"btc24hEarning"`
	Usdt24HEarning   string `json:"usdt24hEarning"`
	BtcTotalEarning  string `json:"btcTotalEarning"`
	UsdtTotalEarning string `json:"usdtTotalEarning"`
}

func (p *SpotClient) SavingsAccount() (*SavingsAccountInfo, error) {
	resp, err := p.DoGet("/api/v2/earn/savings/account", nil)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data *SavingsAccountInfo
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

type SavingsAssets struct {
	ResultList []struct {
		ProductID       string `json:"productId"`
		ProductCoin     string `json:"productCoin"`
		InterestCoin    string `json:"interestCoin"`
		PeriodType      string `json:"periodType"`
		Period          string `json:"period"`
		HoldAmount      string `json:"holdAmount"`
		LastProfit      string `json:"lastProfit"`
		TotalProfit     string `json:"totalProfit"`
		HoldDays        string `json:"holdDays"`
		Status          string `json:"status"`
		AllowRedemption string `json:"allowRedemption"`
		ProductLevel    string `json:"productLevel"`
		Apy             []struct {
			RateLevel  string `json:"rateLevel"`
			MinApy     string `json:"minApy"`
			MaxApy     string `json:"maxApy"`
			CurrentApy string `json:"currentApy"`
		} `json:"apy"`
	} `json:"resultList"`
	EndID string `json:"endId"`
}

// SavingsAssets returns all assets
func (p *SpotClient) SavingsAssets(periodType string) (*SavingsAssets, error) {
	params := map[string]string{
		"periodType": periodType,
	}
	resp, err := p.DoGet("/api/v2/earn/savings/assets", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data *SavingsAssets
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

type SavingsProduct struct {
	ProductId     string `json:"productId"`
	Coin          string `json:"coin"`
	PeriodType    string `json:"periodType"` // flexible/fixed
	Period        string `json:"period"`
	ApyType       string `json:"apyType"`
	AdvanceRedeem string `json:"advanceRedeem"`
	SettleMethod  string `json:"settleMethod"`
	ApyList       []struct {
		RateLevel  string `json:"rateLevel"`
		MinStepVal string `json:"minStepVal"`
		MaxStepVal string `json:"maxStepVal"`
		CurrentApy string `json:"currentApy"`
	} `json:"apyList"`
	Status       string `json:"status"`
	ProductLevel string `json:"productLevel"`
}

// SavingsProduct returns all assets
func (p *SpotClient) SavingsProduct(coin string) ([]*SavingsProduct, error) {
	params := map[string]string{
		"coin": coin,
	}
	resp, err := p.DoGet("/api/v2/earn/savings/product", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*SavingsProduct
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

type SavingsActionResult struct {
	OrderId string `json:"orderId"`
	Status  string `json:"status"`
}

// SavingsSubscribe subscribe to a product, periodType: flexible/fixed
func (p *SpotClient) SavingsSubscribe(productId, periodType string, amount float64) (*SavingsActionResult, error) {
	params := map[string]string{
		"productId":  productId,
		"periodType": periodType,
		"amount":     strconv.FormatFloat(amount, 'f', -1, 64),
	}
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.DoPost("/api/v2/earn/savings/subscribe", postBody)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data *SavingsActionResult
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

// SavingsRedeem redeem from a product, periodType: flexible/fixed
func (p *SpotClient) SavingsRedeem(productId, periodType string, amount float64) (*SavingsActionResult, error) {
	params := map[string]string{
		"productId":  productId,
		"periodType": periodType,
		"amount":     strconv.FormatFloat(amount, 'f', -1, 64),
	}
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.DoPost("/api/v2/earn/savings/redeem", postBody)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data *SavingsActionResult
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}
