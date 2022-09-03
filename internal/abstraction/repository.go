package abstraction

import (
	"gorm.io/gorm"
)

type IRepository interface {
	CheckTrx(ctx *Context) *gorm.DB
}

type Repository struct {
	Connection *gorm.DB
	Db         *gorm.DB
	Tx         *gorm.DB
}

func (r *Repository) CheckTrx(ctx *Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
