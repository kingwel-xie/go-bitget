package model

type SubscribeReq struct {
	InstType string `json:"instType"`
	Channel  string `json:"channel"`
	InstId   string `json:"instId"`
	Coin     string `json:"coin"`
}

func (r SubscribeReq) MakeKey() string {
	return r.InstType + r.Channel
}

func (r SubscribeReq) ToCanonical() SubscribeReq {
	if "" == r.Coin {
		r.Coin = r.InstId
	}
	return r
}
