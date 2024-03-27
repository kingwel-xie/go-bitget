package v2ext

import (
	"encoding/json"
)

type PositionEntry struct {
	MarginCoin       string `json:"marginCoin"`
	Symbol           string `json:"symbol"`
	HoldSide         string `json:"holdSide"`
	OpenDelegateSize string `json:"openDelegateSize"`
	MarginSize       string `json:"marginSize"`
	Available        string `json:"available"`
	Locked           string `json:"locked"`
	Total            string `json:"total"`
	Leverage         string `json:"leverage"`
	AchievedProfits  string `json:"achievedProfits"`
	OpenPriceAvg     string `json:"openPriceAvg"`
	MarginMode       string `json:"marginMode"`
	PosMode          string `json:"posMode"`
	UnrealizedPL     string `json:"unrealizedPL"`
	LiquidationPrice string `json:"liquidationPrice"`
	KeepMarginRate   string `json:"keepMarginRate"`
	MarkPrice        string `json:"markPrice"`
	MarginRatio      string `json:"marginRatio"`
	CTime            string `json:"cTime"`
}

type AccountEntry struct {
	MarginCoin           string `json:"marginCoin"`
	Locked               string `json:"locked"`
	Available            string `json:"available"`
	CrossedMaxAvailable  string `json:"crossedMaxAvailable"`
	IsolatedMaxAvailable string `json:"isolatedMaxAvailable"`
	MaxTransferOut       string `json:"maxTransferOut"`
	AccountEquity        string `json:"accountEquity"`
	UsdtEquity           string `json:"usdtEquity"`
	BtcEquity            string `json:"btcEquity"`
	CrossedRiskRate      string `json:"crossedRiskRate"`
	UnrealizedPL         string `json:"unrealizedPL"`
	Coupon               string `json:"coupon"`
	CrossedUnrealizedPL  string `json:"crossedUnrealizedPL"`
	IsolatedUnrealizedPL string `json:"isolatedUnrealizedPL"`
}

func (p *MixClient) Accounts(productType string) ([]*AccountEntry, error) {
	params := map[string]string{
		"productType": productType,
	}
	resp, err := p.DoGet("/api/v2/mix/account/accounts", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*AccountEntry
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

func (p *MixClient) AllPosition(productType string, marginCoin string) ([]*PositionEntry, error) {
	params := map[string]string{
		"productType": productType,
		"marginCoin":  marginCoin,
	}
	resp, err := p.DoGet("/api/v2/mix/position/all-position", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*PositionEntry
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}
