package course

import (
	"log"

	"github.com/gorilla/websocket"

	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/internal/model/base"
	xws "github.com/mcholismalik/boilerplate-golang/pkg/ws"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	Factory factory.Factory
}

func NewDelivery(f factory.Factory) *delivery {
	return &delivery{f}
}

func (h *delivery) Route(g *echo.Group) {
	g.GET("", h.Course)
}

func (h *delivery) Course(c echo.Context) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	ch := base.NewWsChannel(ws)
	err = xws.HubAssignor(c, ch, h.Factory)
	if err != nil {
		return err
	}

	go xws.Receiver(ch, h.Factory)
	go xws.Sender(ch, h.Factory)
	go ProcessCourse(ch)

	<-ch.Done
	logrus.Println("NewWs: done")
	return nil
}

func ProcessCourse(ch *base.WsChannel) {
	for {
		select {
		case v := <-ch.MsgReceive:
			// if we wanna mask something
			// msg := []byte("drawName('malik')")

			ch.MsgSend <- base.WsMsg{
				MsgType: v.MsgType,
				Msg:     v.Msg,
			}
		case <-ch.Done:
			return
		}
	}
}
