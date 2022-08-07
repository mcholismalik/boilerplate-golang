package ws

import (
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/internal/handler/ws/course"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "ws"
	course.NewHandler(f).Route(e.Group(prefix + "/course"))
}
