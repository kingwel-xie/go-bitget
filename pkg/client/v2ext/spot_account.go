package v2ext

import (
	"encoding/json"
)

type AccountInfo struct {
	UserID      string   `json:"userId"`
	InviterID   string   `json:"inviterId"`
	ChannelCode string   `json:"channelCode"`
	Channel     string   `json:"channel"`
	Ips         string   `json:"ips"`
	Authorities []string `json:"authorities"`
	ParentID    int64    `json:"parentId"`
	TraderType  string   `json:"traderType"`
	RegisTime   int64    `json:"regisTime,string"`
}

func (p *SpotClient) Info() (*AccountInfo, error) {
	resp, err := p.DoGet("/api/v2/spot/account/info", nil)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data *AccountInfo
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}

type AssetEntry struct {
	Coin           string `json:"coin"`
	Available      string `json:"available"`
	LimitAvailable string `json:"limitAvailable"`
	Frozen         string `json:"frozen"`
	Locked         string `json:"locked"`
	UTime          int64  `json:"uTime,string"`
}

// Assets returns all assets, assetType: hold_only/all
func (p *SpotClient) Assets(assetType string) ([]*AssetEntry, error) {
	if assetType == "" {
		assetType = "hold_only"
	}
	params := map[string]string{
		"assetType": assetType,
	}
	resp, err := p.DoGet("/api/v2/spot/account/assets", params)
	if err != nil {
		return nil, err
	}

	var temp struct {
		Response
		Data []*AssetEntry
	}
	err = json.Unmarshal([]byte(resp), &temp)
	if err != nil {
		return nil, err
	}
	return temp.Data, err
}
