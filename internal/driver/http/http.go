package http

import (
	"fmt"
	"net/http"
	"os"

	docs "github.com/mcholismalik/boilerplate-golang/docs"
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/middleware"
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/rest"
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/web"
	"github.com/mcholismalik/boilerplate-golang/internal/delivery/ws"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(f factory.Factory) {
	var (
		APP     = os.Getenv(constant.APP)
		VERSION = os.Getenv(constant.VERSION)
		HOST    = os.Getenv(constant.HOST)
		SCHEME  = os.Getenv(constant.SCHEME)
		PORT    = os.Getenv(constant.PORT)
	)

	// echo
	e := echo.New()

	// middleware
	middleware.Init(e)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// swagger
	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = HOST
	docs.SwaggerInfo.Schemes = []string{SCHEME}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// delivery
	rest.Init(e, f)
	web.Init(e, f)
	ws.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
