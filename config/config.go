package config

import (
	"github.com/kingwel-xie/go-bitget/constants"
)

var (
	BaseUrl = "https://api.bitget.com"
	WsUrl   = "wss://ws.bitget.com/mix/v1/stream"

	ApiKey        = "bg_b85913f93946232babbc8badba3f1d60"
	SecretKey     = "843c074e89dce1cef8acde057304711de462d904bfb8f0c50edf56246688c00b"
	PASSPHRASE    = "88888888"
	TimeoutSecond = 30
	SignType      = constants.SHA256
)

func Init(apiKey, secretKey, passphrase string) {
	ApiKey = apiKey
	SecretKey = secretKey
	PASSPHRASE = passphrase
}
