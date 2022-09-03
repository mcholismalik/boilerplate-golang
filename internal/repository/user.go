package repository

import (
	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	"github.com/mcholismalik/boilerplate-golang/internal/model"

	"gorm.io/gorm"
)

type (
	User interface {
		Base[model.UserModel]
		FindByEmail(ctx abstraction.Context, email string) (*model.UserModel, error)
	}

	user struct {
		Base[model.UserModel]
	}
)

func NewUser(conn *gorm.DB) User {
	model := model.UserModel{}
	base := NewBase(conn, model, model.TableName())
	return &user{
		base,
	}
}

func (m *user) FindByEmail(ctx abstraction.Context, email string) (*model.UserModel, error) {
	m.CheckTrx(ctx)
	query := m.GetConn().Model(model.UserModel{})
	result := new(model.UserModel)
	err := query.Where("email", email).First(result).Error
	if err != nil {
		return nil, m.MaskError(err)
	}
	return result, nil
}
