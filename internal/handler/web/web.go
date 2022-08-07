package web

import (
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/internal/handler/web/bubble"
	"github.com/mcholismalik/boilerplate-golang/internal/handler/web/playground"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "web"
	bubble.NewHandler(f).Route(e.Group(prefix + "/bubble"))
	playground.NewHandler(f).Route(e.Group(prefix + "/playground"))
}
