package v2ext

import (
	"encoding/json"
)

func (p *SpotClient) Symbols(symbol string) (any, error) {
	params := map[string]string{
		"symbol": symbol,
	}
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/symbols", params)
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
