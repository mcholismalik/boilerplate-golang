package migration

import (
	"fmt"

	dbDriver "github.com/mcholismalik/boilerplate-golang/internal/driver/db"
	model "github.com/mcholismalik/boilerplate-golang/internal/model/entity"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/env"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	if !env.NewEnv().GetBool(constant.IS_RUN_MIGRATION) {
		return
	}

	mgConfigurations := map[string]Migration{
		constant.DB_GOBOIL_CLEAN: &migration{
			DbModels: &[]interface{}{
				&model.UserModel{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := dbDriver.GetConnection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			logrus.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}
}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}
