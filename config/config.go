package config

import (
	"github.com/kingwel-xie/go-bitget/constants"
)

var (
	BaseUrl      = "https://api.bitget.com"
	WsUrl        = "wss://ws.bitget.com/v2/ws/public"
	WsPrivateUrl = "wss://ws.bitget.com/v2/ws/private"

	ApiKey        = "bg_b85913f93946232babbc8badba3f1d60"
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
