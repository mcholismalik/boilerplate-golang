package api

import (
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/api/auth"
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/api/user"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	// routes
	prefix := "api"
	auth.NewDelivery(f).Route(e.Group(prefix + "/auth"))
	user.NewDelivery(f).Route(e.Group(prefix + "/users"))
}
