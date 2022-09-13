package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	xtemplate "github.com/mcholismalik/boilerplate-golang/pkg/util/template"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/validator"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	var (
		APP  = os.Getenv(constant.APP)
		ENV  = os.Getenv(constant.ENV)
		NAME = fmt.Sprintf("%s-%s", APP, ENV)
	)

	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
			Format:           fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri}", NAME),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)
	e.HTTPErrorHandler = ErrorHandler
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}
	e.Renderer = xtemplate.NewRenderer("internal/views/*.html", true)
}
