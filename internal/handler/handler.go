package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/internal/handler/rest"
	"github.com/mcholismalik/boilerplate-golang/internal/handler/web"
	"github.com/mcholismalik/boilerplate-golang/internal/handler/ws"
)

func Init(e *echo.Echo, f factory.Factory) {
	rest.Init(e, f)
	web.Init(e, f)
	ws.Init(e, f)
}
