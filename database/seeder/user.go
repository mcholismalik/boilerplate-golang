package seeder

import (
	"log"

	"github.com/mcholismalik/boilerplate-golang/internal/model"

	"gorm.io/gorm"
)

func UserTableSeeder(conn *gorm.DB) {
	trx := conn.Begin()

	if err := trx.Create(&model.UserEntityModel{
		UserEntity: model.UserEntity{
			Name:     "admin",
			Email:    "admin@gmail.com",
			Password: "admin",
		},
		Context: Context,
	}).Error; err != nil {
		trx.Rollback()
		log.Println(err.Error())
		return
	}

	if err := trx.Commit().Error; err != nil {
		log.Println(err.Error())
	}
}
