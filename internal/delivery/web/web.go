package web

import (
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/web/bubble"
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/web/playground"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "web"
	bubble.NewDelivery(f).Route(e.Group(prefix + "/bubble"))
	playground.NewDelivery(f).Route(e.Group(prefix + "/playground"))
}
