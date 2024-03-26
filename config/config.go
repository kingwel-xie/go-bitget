package config

import "github.com/kingwel-xie/go-bitget/constants"

var (
	BaseUrl = "https://api.bitget.com"
	WsUrl   = "wss://ws.bitget.com/mix/v1/stream"

	ApiKey        = ""
	SecretKey     = ""
	PASSPHRASE    = ""
	TimeoutSecond = 30
	SignType      = constants.SHA256
)

func Init(apiKey, secretKey, passphrase string) {
	ApiKey = apiKey
	SecretKey = secretKey
	PASSPHRASE = passphrase
}
