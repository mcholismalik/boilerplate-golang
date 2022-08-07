package main

import (
	"os"

	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/internal/handler/web"
	"github.com/mcholismalik/boilerplate-golang/internal/handler/ws"
	"github.com/mcholismalik/boilerplate-golang/internal/middleware"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/env"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv(constant.ENV)
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title boilerplate-golang
// @version 0.0.1
// @description This is a doc for boilerplate-golang.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {
	var PORT = os.Getenv(constant.PORT)

	// dependency
	// db.Init()
	// migration.Init()
	// seeder.Init()

	e := echo.New()
	middleware.Init(e)

	// factory
	f := factory.Init()

	// handler
	// rest.Init(e, f)
	web.Init(e, f)
	ws.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
