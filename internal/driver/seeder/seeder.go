package seeder

import (
	"context"

	dbDriver "github.com/mcholismalik/boilerplate-golang/internal/driver/db"
	"github.com/mcholismalik/boilerplate-golang/internal/model/abstraction"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	"github.com/mcholismalik/boilerplate-golang/pkg/ctxval"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/env"

	"github.com/google/uuid"
)

var Context = ctxval.SetAuthValue(
	context.Background(),
	&abstraction.AuthContext{
		ID:    uuid.New(),
		Name:  "system",
		Email: "system@system.sys",
	})

func Init() {
	if !env.NewEnv().GetBool(constant.IS_RUN_SEEDER) {
		return
	}

	conn, err := dbDriver.GetConnection(constant.DB_GOBOIL_CLEAN)
	if err != nil {
		panic(err)
	}

	UserTableSeeder(conn)
}
