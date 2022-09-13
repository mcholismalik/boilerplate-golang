package abstraction

import "github.com/gorilla/websocket"

var MsgTypeText = 1

type WsChannel struct {
	User       string
	Ws         *websocket.Conn
	Done       chan struct{}
	MsgReceive chan WsMsg
	MsgSend    chan WsMsg
}

type WsMsg struct {
	MsgType int
	Msg     []byte
}

func NewWsChannel(ws *websocket.Conn) *WsChannel {
	ch := &WsChannel{
		Ws:         ws,
		Done:       make(chan struct{}),
		MsgReceive: make(chan WsMsg),
		MsgSend:    make(chan WsMsg),
	}
	return ch
}

type WsHub struct {
	ChannelSenders   map[string]*WsChannel
	ChannelReceivers map[string]*WsChannel
}

func NewWsHub() *WsHub {
	return &WsHub{
		ChannelSenders:   make(map[string]*WsChannel),
		ChannelReceivers: make(map[string]*WsChannel),
	}
}
