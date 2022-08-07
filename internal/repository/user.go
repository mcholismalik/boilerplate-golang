package repository

import (
	"strings"

	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	"github.com/mcholismalik/boilerplate-golang/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User interface {
	checkTrx(ctx *abstraction.Context) *gorm.DB
	FindAll(ctx *abstraction.Context, payload *abstraction.SearchGetRequest, p *abstraction.Pagination) ([]model.UserEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id uuid.UUID) (model.UserEntityModel, error)
	FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error)
	Create(ctx *abstraction.Context, data model.UserEntityModel) (model.UserEntityModel, error)
	UpdateByID(ctx *abstraction.Context, ID uuid.UUID, data model.UserEntityModel) (model.UserEntityModel, error)
	DeleteByID(ctx *abstraction.Context, id uuid.UUID) error
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *user) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db.WithContext(ctx.Context)
}

func (r *user) FindAll(ctx *abstraction.Context, payload *abstraction.SearchGetRequest, p *abstraction.Pagination) ([]model.UserEntityModel, *abstraction.PaginationInfo, error) {
	var users []model.UserEntityModel
	var count int64

	query := r.checkTrx(ctx).Model(&model.UserEntityModel{})
	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? or lower(email) Like ?", search, search)
	}
	BuildFilterSortQuery[model.UserEntity](model.UserEntity{}, "users", query, payload)

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := abstraction.GetLimitOffset(p)
	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, abstraction.BuildPaginationInfo[model.UserEntityModel](p, count), err
}

func (r *user) FindByID(ctx *abstraction.Context, id uuid.UUID) (model.UserEntityModel, error) {
	return FindByID[model.UserEntityModel](r.CheckTrx(ctx), id, model.UserEntityModel{})
}

func (r *user) FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Create(ctx *abstraction.Context, user model.UserEntityModel) (model.UserEntityModel, error) {
	return Create[model.UserEntityModel](r.checkTrx(ctx), user)
}

func (r *user) UpdateByID(ctx *abstraction.Context, ID uuid.UUID, data model.UserEntityModel) (model.UserEntityModel, error) {
	return UpdateByID[model.UserEntityModel](r.checkTrx(ctx), ID, data)
}

func (r *user) DeleteByID(ctx *abstraction.Context, id uuid.UUID) error {
	return DeleteByID[model.UserEntityModel](r.CheckTrx(ctx), id, model.UserEntityModel{})
}
