package rest

import (
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/rest/auth"
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/rest/user"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	// routes
	prefix := "rest"
	auth.NewDelivery(f).Route(e.Group(prefix + "/auth"))
	user.NewDelivery(f).Route(e.Group(prefix + "/users"))
}
