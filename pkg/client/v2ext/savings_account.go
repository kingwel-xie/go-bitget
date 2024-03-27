package v2ext

import (
	"encoding/json"
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

// SavingsAssets returns all assets, periodType: flexible/fixed
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
