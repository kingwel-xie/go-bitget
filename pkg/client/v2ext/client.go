package v2ext

import (
	"github.com/kingwel-xie/go-bitget/internal/common"
)

type MixClient struct {
	*common.BitgetRestClient
}

func NewMixClient() *MixClient {
	return &MixClient{
		new(common.BitgetRestClient).Init(),
	}
}

type SpotClient struct {
	*common.BitgetRestClient
}

func NewSpotClient() *SpotClient {
	return &SpotClient{
		new(common.BitgetRestClient).Init(),
	}
}

type Response struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}
