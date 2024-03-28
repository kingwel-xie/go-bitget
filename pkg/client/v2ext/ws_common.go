package v2ext

type GenericMessage struct {
	Action string `json:"action"`
	Arg    struct {
		InstType string `json:"instType"`
		Channel  string `json:"channel"`
		InstID   string `json:"instId"`
	} `json:"arg"`
	Ts int64 `json:"ts"`
}

// WsHandler handle raw websocket message
type WsHandler func(string)

// ErrHandler handles errors
type ErrHandler func(error)
