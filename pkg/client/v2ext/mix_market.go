package v2ext

import (
	"encoding/json"
	"strconv"
)

type ContractItem struct {
	Symbol              string   `json:"symbol"`
	BaseCoin            string   `json:"baseCoin"`
	QuoteCoin           string   `json:"quoteCoin"`
	BuyLimitPriceRatio  string   `json:"buyLimitPriceRatio"`
	SellLimitPriceRatio string   `json:"sellLimitPriceRatio"`
	FeeRateUpRatio      string   `json:"feeRateUpRatio"`
	MakerFeeRate        string   `json:"makerFeeRate"`
	TakerFeeRate        string   `json:"takerFeeRate"`
	OpenCostUpRatio     string   `json:"openCostUpRatio"`
	SupportMarginCoins  []string `json:"supportMarginCoins"`
	MinTradeNum         string   `json:"minTradeNum"`
	PriceEndStep        string   `json:"priceEndStep"`
	VolumePlace         string   `json:"volumePlace"`
	PricePlace          string   `json:"pricePlace"`
	SizeMultiplier      string   `json:"sizeMultiplier"`
	SymbolType          string   `json:"symbolType"`
	MinTradeUSDT        string   `json:"minTradeUSDT"`
	MaxSymbolOrderNum   string   `json:"maxSymbolOrderNum"`
	MaxProductOrderNum  string   `json:"maxProductOrderNum"`
	MaxPositionNum      string   `json:"maxPositionNum"`
	SymbolStatus        string   `json:"symbolStatus"`
	OffTime             string   `json:"offTime"`
	LimitOpenTime       string   `json:"limitOpenTime"`
	DeliveryTime        string   `json:"deliveryTime"`
	DeliveryStartTime   string   `json:"deliveryStartTime"`
	DeliveryPeriod      string   `json:"deliveryPeriod"`
	LaunchTime          string   `json:"launchTime"`
	FundInterval        string   `json:"fundInterval"`
	MinLever            string   `json:"minLever"`
	MaxLever            string   `json:"maxLever"`
	PosLimit            string   `json:"posLimit"`
	MaintainTime        string   `json:"maintainTime"`
}

func (p *MixClient) Contracts(productType string, args ...string) ([]*ContractItem, error) {
	params := map[string]string{
		"productType": productType,
	}
	if len(args) > 0 {
		params["symbol"] = args[0]
	}

	resp, err := p.DoGet("/api/v2/mix/market/contracts", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*ContractItem
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

type FundingRateItem struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime,string"`
}

func (p *MixClient) HistoryFundingRate(productType string, symbol string, pageNo int, pageSize int) ([]*FundingRateItem, error) {
	params := map[string]string{
		"productType": productType,
		"symbol":      symbol,
		"pageNo":      strconv.Itoa(pageNo),
		"pageSize":    strconv.Itoa(pageSize),
	}

	resp, err := p.DoGet("/api/v2/mix/market/history-fund-rate", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*FundingRateItem
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

func (p *MixClient) CurrentFundingRate(productType string, symbol string) ([]*FundingRateItem, error) {
	params := map[string]string{
		"productType": productType,
		"symbol":      symbol,
	}

	resp, err := p.DoGet("/api/v2/mix/market/current-fund-rate", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*FundingRateItem
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

type TickerEntry struct {
	Symbol            string      `json:"symbol"`
	LastPr            string      `json:"lastPr"`
	AskPr             string      `json:"askPr"`
	BidPr             string      `json:"bidPr"`
	BidSz             string      `json:"bidSz"`
	AskSz             string      `json:"askSz"`
	High24H           string      `json:"high24h"`
	Low24H            string      `json:"low24h"`
	Ts                int64       `json:"ts,string"`
	Change24H         string      `json:"change24h"`
	BaseVolume        string      `json:"baseVolume"`
	QuoteVolume       string      `json:"quoteVolume"`
	UsdtVolume        string      `json:"usdtVolume"`
	OpenUtc           string      `json:"openUtc"`
	ChangeUtc24H      string      `json:"changeUtc24h"`
	IndexPrice        string      `json:"indexPrice"`
	FundingRate       string      `json:"fundingRate"`
	HoldingAmount     string      `json:"holdingAmount"`
	DeliveryStartTime interface{} `json:"deliveryStartTime"`
	DeliveryTime      interface{} `json:"deliveryTime"`
	DeliveryStatus    string      `json:"deliveryStatus"`
	Open24H           string      `json:"open24h"`
}

func (p *MixClient) Ticker(productType string, symbol string) ([]*TickerEntry, error) {
	params := map[string]string{
		"productType": productType,
		"symbol":      symbol,
	}
	resp, err := p.DoGet("/api/v2/mix/market/ticker", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*TickerEntry
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, nil
}

func (p *MixClient) Tickers(productType string) ([]*TickerEntry, error) {
	params := map[string]string{
		"productType": productType,
	}

	resp, err := p.DoGet("/api/v2/mix/market/tickers", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*TickerEntry
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, nil
}
