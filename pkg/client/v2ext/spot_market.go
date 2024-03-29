package v2ext

import (
	"encoding/json"
	"github.com/kingwel-xie/go-bitget/internal"
	"strconv"
)

type CoinInfo struct {
	CoinId   string `json:"coinId"`
	Coin     string `json:"coin"`
	Transfer string `json:"transfer"`
	Chains   []struct {
		Chain             string `json:"chain"`
		NeedTag           string `json:"needTag"`
		Withdrawable      string `json:"withdrawable"`
		Rechargeable      string `json:"rechargeable"`
		WithdrawFee       string `json:"withdrawFee"`
		ExtraWithdrawFee  string `json:"extraWithdrawFee"`
		DepositConfirm    string `json:"depositConfirm"`
		WithdrawConfirm   string `json:"withdrawConfirm"`
		MinDepositAmount  string `json:"minDepositAmount"`
		MinWithdrawAmount string `json:"minWithdrawAmount"`
		BrowserUrl        string `json:"browserUrl"`
		ContractAddress   string `json:"contractAddress"`
		WithdrawStep      string `json:"withdrawStep"`
	} `json:"chains"`
}

func (p *SpotClient) Coins(coin string) ([]*CoinInfo, error) {
	params := map[string]string{}
	if coin != "" {
		params["coin"] = coin
	}
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/coins", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*CoinInfo
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, nil
}

type SymbolInfo struct {
	Symbol              string `json:"symbol"`
	BaseCoin            string `json:"baseCoin"`
	QuoteCoin           string `json:"quoteCoin"`
	MinTradeAmount      string `json:"minTradeAmount"`
	MaxTradeAmount      string `json:"maxTradeAmount"`
	TakerFeeRate        string `json:"takerFeeRate"`
	MakerFeeRate        string `json:"makerFeeRate"`
	PricePrecision      string `json:"pricePrecision"`
	QuantityPrecision   string `json:"quantityPrecision"`
	QuotePrecision      string `json:"quotePrecision"`
	Status              string `json:"status"`
	MinTradeUSDT        string `json:"minTradeUSDT"`
	BuyLimitPriceRatio  string `json:"buyLimitPriceRatio"`
	SellLimitPriceRatio string `json:"sellLimitPriceRatio"`
}

func (p *SpotClient) Symbols(symbol string) ([]*SymbolInfo, error) {
	params := map[string]string{}
	if symbol != "" {
		params["symbol"] = symbol
	}
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/symbols", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*SymbolInfo
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, nil
}

type SpotTickerEntry struct {
	Symbol       string `json:"symbol"`
	High24H      string `json:"high24h"`
	Open         string `json:"open"`
	Low24H       string `json:"low24h"`
	LastPr       string `json:"lastPr"`
	QuoteVolume  string `json:"quoteVolume"`
	BaseVolume   string `json:"baseVolume"`
	UsdtVolume   string `json:"usdtVolume"`
	BidPr        string `json:"bidPr"`
	AskPr        string `json:"askPr"`
	BidSz        string `json:"bidSz"`
	AskSz        string `json:"askSz"`
	OpenUtc      string `json:"openUtc"`
	Ts           int64  `json:"ts,string"`
	ChangeUtc24H string `json:"changeUtc24h"`
	Change24H    string `json:"change24h"`
}

func (p *SpotClient) Tickers(symbol string) ([]*SpotTickerEntry, error) {
	params := map[string]string{
		"symbol": symbol,
	}
	resp, err := p.DoGet("/api/v2/spot/market/tickers", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*SpotTickerEntry
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, nil
}

type CandleItem struct {
	Ts          int64
	OpenPrice   float64
	HighPrice   float64
	LowPrice    float64
	ClosePrice  float64
	BaseVolume  float64
	QuoteVolume float64
	UsdtVolume  float64
}

func (p *SpotClient) Candles(symbol, granularity string, startTime, endTime int64, limit int) ([]*CandleItem, error) {
	params := map[string]string{
		"symbol":      symbol,
		"granularity": granularity,
		"startTime":   strconv.FormatInt(startTime, 10),
		"endTime":     strconv.FormatInt(endTime, 10),
		"limit":       strconv.Itoa(limit),
	}
	if startTime > 0 {
		params["startTime"] = strconv.FormatInt(startTime, 10)
	}
	if endTime > 0 {
		params["endTime"] = strconv.FormatInt(endTime, 10)
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	resp, err := p.DoGet("/api/v2/spot/market/candles", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data [][]string `json:"data"`
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}

	var list []*CandleItem
	for _, i := range temp.Data {
		if len(i) != 7 {
			// skip invalid data
			continue
		}
		list = append(list, &CandleItem{
			Ts:          internal.String2Int64(i[0]),
			OpenPrice:   internal.String2Float(i[1]),
			HighPrice:   internal.String2Float(i[2]),
			LowPrice:    internal.String2Float(i[3]),
			ClosePrice:  internal.String2Float(i[4]),
			BaseVolume:  internal.String2Float(i[5]),
			QuoteVolume: internal.String2Float(i[6]),
			UsdtVolume:  internal.String2Float(i[7]),
		})
	}
	return list, nil
}
