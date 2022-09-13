package main

import (
	"os"

	"github.com/mcholismalik/boilerplate-golang/internal/driver/db"
	"github.com/mcholismalik/boilerplate-golang/internal/driver/http"
	"github.com/mcholismalik/boilerplate-golang/internal/driver/migration"
	"github.com/mcholismalik/boilerplate-golang/internal/driver/seeder"
	"github.com/mcholismalik/boilerplate-golang/internal/driver/sentry"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/env"

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
	// repository driver
	db.Init()
	migration.Init()
	seeder.Init()

	// factory
	f := factory.Init()

	// delivery driver
	sentry.Init()
	http.Init(f)
}
