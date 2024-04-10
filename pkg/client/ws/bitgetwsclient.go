package ws

import (
	"github.com/kingwel-xie/go-bitget/constants"
	"github.com/kingwel-xie/go-bitget/internal/common"
	"github.com/kingwel-xie/go-bitget/internal/model"
	"github.com/kingwel-xie/go-bitget/logging/applogger"
)

type BitgetWsClient struct {
	bitgetBaseWsClient *common.BitgetBaseWsClient
	NeedLogin          bool
}

func (p *BitgetWsClient) Init(needLogin bool, listener common.OnReceive, errorListener common.OnError) (*BitgetWsClient, chan struct{}, chan struct{}, error) {
	p.bitgetBaseWsClient = new(common.BitgetBaseWsClient).Init(needLogin)
	p.bitgetBaseWsClient.SetListener(listener, errorListener)
	err := p.bitgetBaseWsClient.ConnectWebSocket(needLogin)
	if err != nil {
		return nil, nil, nil, err
	}
	doneCh, ctrlCh := p.bitgetBaseWsClient.StartReadLoop()
	p.bitgetBaseWsClient.ExecuterPing()

	if needLogin {
		applogger.Info("login in ...")
		p.bitgetBaseWsClient.Login()
		for {
			if !p.bitgetBaseWsClient.LoginStatus {
				continue
			}
			break
		}
		applogger.Info("login in ... success")
	}

	return p, doneCh, ctrlCh, nil

}

func (p *BitgetWsClient) Connect() *BitgetWsClient {
	p.bitgetBaseWsClient.Connect()
	return p
}

func (p *BitgetWsClient) UnSubscribe(list []model.SubscribeReq) {

	var args []interface{}
	for _, req := range list {
		delete(p.bitgetBaseWsClient.ScribeMap, req.MakeKey())
		p.bitgetBaseWsClient.AllSuribe.Add(req)
		p.bitgetBaseWsClient.AllSuribe.Remove(req)
		args = append(args, req)
	}

	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpUnsubscribe,
		Args: args,
	}

	p.SendMessageByType(wsBaseReq)
}

func (p *BitgetWsClient) SubscribeDef(list []model.SubscribeReq) {
	var args []interface{}
	for _, req := range list {
		req = req.ToCanonical()
		args = append(args, req)
	}
	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}
	p.SendMessageByType(wsBaseReq)
}

func (p *BitgetWsClient) Subscribe(list []model.SubscribeReq, listener common.OnReceive) {
	var args []interface{}
	for _, req := range list {
		req = req.ToCanonical()
		args = append(args, req)
		p.bitgetBaseWsClient.ScribeMap[req.MakeKey()] = listener
		p.bitgetBaseWsClient.AllSuribe.Add(req)
		args = append(args, req)
	}

	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}

	p.bitgetBaseWsClient.SendByType(wsBaseReq)
}

func (p *BitgetWsClient) SubscribeOne(req model.SubscribeReq, listener common.OnReceive) {
	req = req.ToCanonical()
	p.bitgetBaseWsClient.ScribeMap[req.MakeKey()] = listener
	p.bitgetBaseWsClient.AllSuribe.Add(req)

	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: []interface{}{req},
	}

	p.bitgetBaseWsClient.SendByType(wsBaseReq)
}

func (p *BitgetWsClient) SendMessage(msg string) {
	p.bitgetBaseWsClient.Send(msg)
}

func (p *BitgetWsClient) SendMessageByType(req model.WsBaseReq) {
	p.bitgetBaseWsClient.SendByType(req)
}
