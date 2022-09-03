package repository

import (
	"fmt"
	"math"
	"strings"

	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type (
	Base[T any] interface {
		GetConn() *gorm.DB
		CheckTrx(ctx abstraction.Context)
		MaskError(err error) error

		Find(ctx abstraction.Context, filterParam *abstraction.FilterParam) ([]T, *abstraction.PaginationInfo, error)
		FindByID(ctx abstraction.Context, id string) (*T, error)
		Create(ctx abstraction.Context, data *T) (*T, error)
		Creates(ctx abstraction.Context, data []T) ([]T, error)
		UpdateByID(ctx abstraction.Context, id string, data *T) (*T, error)
		DeleteByID(ctx abstraction.Context, id string) error

		BuildFilterSort(name string, query *gorm.DB, filterParam *abstraction.FilterParam)
		BuildPagination(ctx abstraction.Context, query *gorm.DB, pagination abstraction.Pagination) *abstraction.PaginationInfo
	}

	base[T any] struct {
		conn       *gorm.DB
		entity     T
		entityName string
	}
)

func NewBase[T any](conn *gorm.DB, entity T, entityName string) Base[T] {
	return &base[T]{
		conn,
		entity,
		entityName,
	}
}

func (m *base[T]) GetConn() *gorm.DB {
	return m.conn
}

func (m *base[T]) CheckTrx(ctx abstraction.Context) {
	if ctx.Trx != nil {
		m.conn = ctx.Trx.Db
	}
	m.conn = m.conn.WithContext(ctx)
}

func (m *base[T]) Find(ctx abstraction.Context, filterParam *abstraction.FilterParam) ([]T, *abstraction.PaginationInfo, error) {
	m.CheckTrx(ctx)
	query := m.conn.Model(m.entity)

	m.BuildFilterSort(m.entityName, query, filterParam)
	info := m.BuildPagination(ctx, query, filterParam.Pagination)

	result := []T{}
	err := query.Find(&result).Error

	if err != nil {
		return nil, info, err
	}
	return result, info, nil
}

func (m *base[T]) FindByID(ctx abstraction.Context, id string) (*T, error) {
	m.CheckTrx(ctx)
	query := m.conn.Model(m.entity)
	result := new(T)
	err := query.Where("id", id).First(result).Error
	if err != nil {
		return nil, m.MaskError(err)
	}
	return result, nil
}

func (m *base[T]) Create(ctx abstraction.Context, data *T) (*T, error) {
	m.CheckTrx(ctx)
	query := m.conn.Model(m.entity)
	err := query.Create(data).Error
	return data, m.MaskError(err)
}
func (m *base[T]) Creates(ctx abstraction.Context, data []T) ([]T, error) {
	m.CheckTrx(ctx)
	err := m.conn.Model(m.entity).Create(&data).Error
	return data, m.MaskError(err)
}

func (m *base[T]) UpdateByID(ctx abstraction.Context, id string, data *T) (*T, error) {
	m.CheckTrx(ctx)
	err := m.conn.Model(data).Updates(data).Error
	return data, m.MaskError(err)
}

func (m *base[T]) DeleteByID(ctx abstraction.Context, id string) error {
	m.CheckTrx(ctx)
	err := m.conn.Model(m.entity).Where("id = ?", id).Error
	return m.MaskError(err)
}

func (m *base[T]) BuildFilterSort(name string, query *gorm.DB, filterParam *abstraction.FilterParam) {
	for _, filter := range filterParam.Query {
		query.Where(filter.Field+" = ?", filter.Value)
	}

	for i := range filterParam.SortBy {
		sortBys := strings.Split(filterParam.SortBy[i], ",")
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

func (m *base[T]) BuildPagination(ctx abstraction.Context, tx *gorm.DB, pagination abstraction.Pagination) *abstraction.PaginationInfo {
	info := &abstraction.PaginationInfo{}
	if pagination.Page != nil {
		limit := 10
		if pagination.Limit != nil {
			limit = *pagination.Limit
		}
		page := 0
		if *pagination.Page >= 0 {
			page = *pagination.Page
		}

		tx.Count(&info.Count)
		offset := (page - 1) * limit
		tx.Limit(limit).Offset(offset)
		info.TotalPage = int64(math.Ceil(float64(info.Count) / float64(limit)))

		info.Pagination = pagination
	}

	return info
}

func (m *base[T]) MaskError(err error) error {
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
