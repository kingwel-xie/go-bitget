package common

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kingwel-xie/go-bitget/config"
	"github.com/kingwel-xie/go-bitget/constants"
	"github.com/kingwel-xie/go-bitget/internal"
	"github.com/kingwel-xie/go-bitget/internal/model"
	"github.com/kingwel-xie/go-bitget/logging/applogger"
	"github.com/robfig/cron"
	"sync"
	"time"
)

type BitgetBaseWsClient struct {
	NeedLogin        bool
	Connection       bool
	LoginStatus      bool
	Listener         OnReceive
	ErrorListener    OnError
	Ticker           *time.Ticker
	SendMutex        *sync.Mutex
	WebSocketClient  *websocket.Conn
	LastReceivedTime time.Time
	AllSuribe        *model.Set
	Signer           *Signer
	ScribeMap        map[string]OnReceive
}

func (p *BitgetBaseWsClient) Init(needLogin bool) *BitgetBaseWsClient {
	p.NeedLogin = needLogin
	p.Connection = false
	p.AllSuribe = model.NewSet()
	p.Signer = new(Signer).Init(config.SecretKey)
	p.ScribeMap = make(map[string]OnReceive)
	p.SendMutex = &sync.Mutex{}
	p.Ticker = time.NewTicker(constants.TimerIntervalSecond * time.Second)
	p.LastReceivedTime = time.Now()

	return p
}

func (p *BitgetBaseWsClient) SetListener(msgListener OnReceive, errorListener OnError) {
	p.Listener = msgListener
	p.ErrorListener = errorListener
}

func (p *BitgetBaseWsClient) Connect() {

	p.tickerLoop()
	p.ExecuterPing()
}

func (p *BitgetBaseWsClient) ConnectWebSocket(isPrivate bool) error {
	var err error
	applogger.Info("WebSocket connecting...")
	var url = config.WsUrl
	if isPrivate {
		url = config.WsPrivateUrl
	}
	p.WebSocketClient, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		applogger.Error("WebSocket connected error: %s\n", err)
		return err
	}
	applogger.Info("WebSocket connected")
	p.Connection = true
	return nil
}

func (p *BitgetBaseWsClient) Login() {
	timesStamp := internal.TimesStampSec()
	sign := p.Signer.Sign(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)
	if constants.RSA == config.SignType {
		sign = p.Signer.SignByRSA(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)
	}

	loginReq := model.WsLoginReq{
		ApiKey:     config.ApiKey,
		Passphrase: config.PASSPHRASE,
		Timestamp:  timesStamp,
		Sign:       sign,
	}
	var args []interface{}
	args = append(args, loginReq)

	baseReq := model.WsBaseReq{
		Op:   constants.WsOpLogin,
		Args: args,
	}
	p.SendByType(baseReq)
}

func (p *BitgetBaseWsClient) StartReadLoop() (chan struct{}, chan struct{}) {
	doneCh := make(chan struct{})
	ctrlCh := make(chan struct{})
	go p.ReadLoop(doneCh, ctrlCh)
	return doneCh, ctrlCh
}

func (p *BitgetBaseWsClient) ExecuterPing() {
	c := cron.New()
	_ = c.AddFunc("*/15 * * * * *", p.ping)
	c.Start()
}
func (p *BitgetBaseWsClient) ping() {
	p.Send("ping")
}

func (p *BitgetBaseWsClient) SendByType(req model.WsBaseReq) {
	json, _ := internal.ToJson(req)
	p.Send(json)
}

func (p *BitgetBaseWsClient) Send(data string) {
	if p.WebSocketClient == nil {
		applogger.Error("WebSocket sent error: no connection available")
		return
	}
	applogger.Debug("sendMessage:%s", data)
	p.SendMutex.Lock()
	err := p.WebSocketClient.WriteMessage(websocket.TextMessage, []byte(data))
	p.SendMutex.Unlock()
	if err != nil {
		applogger.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
}

func (p *BitgetBaseWsClient) tickerLoop() {
	applogger.Info("tickerLoop started")
	for {
		select {
		case <-p.Ticker.C:
			elapsedSecond := time.Now().Sub(p.LastReceivedTime).Seconds()

			if elapsedSecond > constants.ReconnectWaitSecond {
				applogger.Info("WebSocket reconnect...")
				p.disconnectWebSocket()
				p.ConnectWebSocket(p.NeedLogin)
			}
		}
	}
}

func (p *BitgetBaseWsClient) disconnectWebSocket() {
	if p.WebSocketClient == nil {
		return
	}

	applogger.Info("WebSocket disconnecting...")
	err := p.WebSocketClient.Close()
	if err != nil {
		applogger.Error("WebSocket disconnect error: %s\n", err)
		return
	}

	applogger.Info("WebSocket disconnected")
}

func (p *BitgetBaseWsClient) ReadLoop(doneCh chan struct{}, ctrlCh chan struct{}) {
	// anyway, either ReadMessage returns error or ctrlCh is closed externally, we have to close doneCh
	defer close(doneCh)

	if p.WebSocketClient == nil {
		applogger.Info("Read error: no connection available")
		//time.Sleep(TimerIntervalSecond * time.Second)
		return
	}

	// Wait for the channel to be closed.  ReadMessage is a blocking operation, so we do
	// it in a separated routine
	silent := false
	go func() {
		select {
		case <-ctrlCh:
			silent = true
		case <-doneCh:
		}
		p.WebSocketClient.Close()
	}()

	for {
		_, buf, err := p.WebSocketClient.ReadMessage()
		if err != nil {
			applogger.Info("Read error: %s", err)
			if !silent {
				p.ErrorListener(err)
			}
			return
		}
		p.LastReceivedTime = time.Now()
		message := string(buf)

		if message == "pong" {
			applogger.Debug("Keep connected:" + message)
			continue
		}
		jsonMap := internal.JSONToMap(message)

		v, e := jsonMap["code"]
		if e && int(v.(float64)) != 0 {
			p.ErrorListener(fmt.Errorf("error, code %v", v))
			continue
		}

		v, e = jsonMap["event"]
		if e && v == "login" {
			applogger.Info("login msg:" + message)
			p.LoginStatus = true
			continue
		}

		v, e = jsonMap["data"]
		if e {
			listener := p.GetListener(jsonMap["arg"])
			listener(message)
			continue
		}
		p.handleMessage(message)
	}

}

func (p *BitgetBaseWsClient) GetListener(argJson interface{}) OnReceive {
	mapData := argJson.(map[string]interface{})
	key := fmt.Sprintf("%s%s", mapData["instType"], mapData["channel"])
	v, e := p.ScribeMap[key]
	if !e {
		return p.Listener
	}
	return v
}

type OnReceive func(string)
type OnError func(error)

func (p *BitgetBaseWsClient) handleMessage(msg string) {
	//fmt.Println("default:" + msg)
}
