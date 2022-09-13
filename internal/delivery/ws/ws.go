package ws

import (
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/ws/course"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "ws"
	course.NewDelivery(f).Route(e.Group(prefix + "/course"))
}
