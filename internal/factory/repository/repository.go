package repository

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/messaging"
	dbDriver "github.com/mcholismalik/boilerplate-golang/internal/driver/db"
	dbRepository "github.com/mcholismalik/boilerplate-golang/internal/repository/db"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	el "github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type Factory struct {
	Db        *gorm.DB
	Es        *el.Client
	Fcm       *messaging.Client
	Firestore *firestore.Client

	UserRepository dbRepository.User
}

func Init() Factory {
	f := Factory{}
	f.InitDb()
	f.InitUserDbRepository()

	return f
}

func (f *Factory) InitDb() {
	db, err := dbDriver.GetConnection(constant.DB_GOBOIL_CLEAN)
	if err != nil {
		panic("Failed init db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) InitUserDbRepository() {
	if f.Db == nil {
		panic("Failed init repository, db is undefined")
	}

	f.UserRepository = dbRepository.NewUser(f.Db)
}
