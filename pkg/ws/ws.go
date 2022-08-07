package ws

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/sirupsen/logrus"
)

func HubAssignor(c echo.Context, ch *abstraction.WsChannel, f factory.Factory) error {
	queries := c.Request().URL.Query()
	allowFields := map[string]bool{
		"sender":   true,
		"receiver": true,
	}
	for field, values := range queries {
		if field != "key" || len(values) == 0 {
			return errors.New("invalid key")
		}

		value := values[0]
		if !strings.Contains(value, ":") {
			return errors.New("invalid key format")
		}

		valueArgs := strings.Split(value, ":")
		if !allowFields[valueArgs[0]] || len(valueArgs) < 2 {
			return errors.New("invalid key format length")
		}

		switch valueArgs[0] {
		case "sender":
			if _, ok := f.WsHub.ChannelSenders[valueArgs[1]]; !ok {
				f.WsHub.ChannelSenders[valueArgs[1]] = ch
			}
		case "receiver":
			if _, ok := f.WsHub.ChannelReceivers[valueArgs[1]]; !ok {
				f.WsHub.ChannelReceivers[valueArgs[1]] = ch
			}
		}

		logrus.Println("WsHubAssignor:", *f.WsHub)

		ch.User = valueArgs[1]
	}

	return nil
}

func Receiver(ch *abstraction.WsChannel, f factory.Factory) {
	defer func() {
		ch.Ws.Close()
	}()

	for {
		msgType, msg, err := ch.Ws.ReadMessage()
		if err != nil {
			break
		}
		logrus.Println("wsReceiver:", msgType, string(msg))
		ch.MsgReceive <- abstraction.WsMsg{
			MsgType: msgType,
			Msg:     msg,
		}
	}
	close(ch.Done)
}

func Sender(ch *abstraction.WsChannel, f factory.Factory) {
	defer func() {
		ch.Ws.Close()
	}()

breakLoop:
	for {
		select {
		case v := <-ch.MsgSend:
			conn := ch.Ws
			val, ok := f.WsHub.ChannelReceivers[ch.User]
			if ok {
				conn = val.Ws
			}

			err := conn.WriteMessage(v.MsgType, v.Msg)
			if err != nil {
				break breakLoop
			}
			logrus.Println("wsSender:", v.MsgType, string(v.Msg))
		case <-ch.Done:
			return
		}
	}
	close(ch.Done)
}
