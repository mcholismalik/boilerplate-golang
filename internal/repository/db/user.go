package db

import (
	"context"

	model "github.com/mcholismalik/boilerplate-golang/internal/model/entity"

	"gorm.io/gorm"
)

type (
	User interface {
		Base[model.UserModel]
		FindByEmail(ctx context.Context, email string) (*model.UserModel, error)
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

func (m *user) FindByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	m.checkTrx(ctx)
	query := m.getConn().Model(model.UserModel{})
	result := new(model.UserModel)
	err := query.Where("email", email).First(result).Error
	if err != nil {
		return nil, m.maskError(err)
	}
	return result, nil
}
