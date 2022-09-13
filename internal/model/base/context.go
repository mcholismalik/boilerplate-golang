package base

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthContext struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type TrxContext struct {
	Db *gorm.DB
}
