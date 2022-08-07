package abstraction

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Context struct {
	context.Context
	Auth *AuthContext
	Trx  *TrxContext
}

type AuthContext struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type TrxContext struct {
	Db *gorm.DB
}
