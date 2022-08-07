package repository

import (
	"fmt"
	"strings"

	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	"github.com/mcholismalik/boilerplate-golang/internal/model"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type validType interface {
	model.UserEntityModel | model.UserEntity
}

func FindAll[T validType](conn *gorm.DB, data T) (T, error) {
	err := conn.Find(&data).Error
	return data, maskError(err)
}

func FindByID[T validType](conn *gorm.DB, ID uuid.UUID, data T) (T, error) {
	err := conn.Where("id = ?", ID).First(&data).Error
	return data, maskError(err)
}

func Create[T validType](conn *gorm.DB, data T) (T, error) {
	err := conn.Create(&data).Error
	return data, maskError(err)
}

func UpdateByID[T validType](conn *gorm.DB, ID uuid.UUID, data T) (T, error) {
	err := conn.Model(&data).Where("id = ?", ID).Updates(&data).Error
	return data, maskError(err)
}

func DeleteByID[T validType](conn *gorm.DB, ID uuid.UUID, data T) error {
	return maskError(conn.Where("id = ?", ID).Delete(&data).Error)
}

func Deletes[T validType](conn *gorm.DB, IDs []uuid.UUID, data T) error {
	return maskError(conn.Delete(&data, IDs).Error)
}

func BuildFilterSortQuery[T validType](data T, name string, query *gorm.DB, payload *abstraction.SearchGetRequest) {
	for _, filter := range payload.Filters {
		query.Where(filter.Field+" = ?", filter.Value)
	}

	for i := range payload.SortBy {
		sortBys := strings.Split(payload.SortBy[i], ",")
		for _, sortBy := range sortBys {
			prefix := sortBy[:1]
			sortType := "asc"
			if prefix == "-" {
				sortType = "desc"
				sortBy = sortBy[1:]
			}

			sortArg := fmt.Sprintf("%s.%s %s", name, sortBy, sortType)
			query = query.Order(sortArg)
		}
	}
}

func maskError(err error) error {
	if err != nil {
		// not found
		if err == gorm.ErrRecordNotFound {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		if pqErr, ok := err.(*pgconn.PgError); ok {
			// duplicate data
			if pqErr.Code == "23505" {
				return res.ErrorBuilder(&res.ErrorConstant.DuplicateEntity, err)
			}
		}

		return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
	}

	return nil
}
