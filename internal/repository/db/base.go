package db

import (
	"context"
	"fmt"
	"math"
	"strings"

	abstraction "github.com/mcholismalik/boilerplate-golang/internal/model/base"
	"github.com/mcholismalik/boilerplate-golang/pkg/ctxval"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type (
	Base[T any] interface {
		getConn() *gorm.DB
		checkTrx(ctx context.Context)
		maskError(err error) error
		buildFilterSort(name string, query *gorm.DB, filterParam abstraction.Filter)
		buildPagination(ctx context.Context, query *gorm.DB, pagination abstraction.Pagination) *abstraction.PaginationInfo

		Find(ctx context.Context, filterParam abstraction.Filter) ([]T, *abstraction.PaginationInfo, error)
		FindByID(ctx context.Context, id string) (*T, error)
		Create(ctx context.Context, data T) (T, error)
		Creates(ctx context.Context, data []T) ([]T, error)
		UpdateByID(ctx context.Context, id string, data T) (T, error)
		DeleteByID(ctx context.Context, id string) error
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

func (m *base[T]) getConn() *gorm.DB {
	return m.conn
}

func (m *base[T]) checkTrx(ctx context.Context) {
	trx := ctxval.GetTrxValue(ctx)
	if trx != nil {
		m.conn = trx.Db
	}
	m.conn = m.conn.WithContext(ctx)
}

func (m *base[T]) buildFilterSort(name string, query *gorm.DB, filterParam abstraction.Filter) {
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

func (m *base[T]) buildPagination(ctx context.Context, tx *gorm.DB, pagination abstraction.Pagination) *abstraction.PaginationInfo {
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

func (m *base[T]) maskError(err error) error {
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

func (m *base[T]) Find(ctx context.Context, filterParam abstraction.Filter) ([]T, *abstraction.PaginationInfo, error) {
	m.checkTrx(ctx)
	query := m.conn.Model(m.entity)

	m.buildFilterSort(m.entityName, query, filterParam)
	info := m.buildPagination(ctx, query, filterParam.Pagination)

	result := []T{}
	err := query.Find(&result).Error

	if err != nil {
		return nil, info, err
	}
	return result, info, nil
}

func (m *base[T]) FindByID(ctx context.Context, id string) (*T, error) {
	m.checkTrx(ctx)
	query := m.conn.Model(m.entity)
	result := new(T)
	err := query.Where("id", id).First(result).Error
	if err != nil {
		return nil, m.maskError(err)
	}
	return result, nil
}

func (m *base[T]) Create(ctx context.Context, data T) (T, error) {
	m.checkTrx(ctx)
	query := m.conn.Model(m.entity)
	err := query.Create(data).Error
	return data, m.maskError(err)
}
func (m *base[T]) Creates(ctx context.Context, data []T) ([]T, error) {
	m.checkTrx(ctx)
	err := m.conn.Model(m.entity).Create(&data).Error
	return data, m.maskError(err)
}

func (m *base[T]) UpdateByID(ctx context.Context, id string, data T) (T, error) {
	m.checkTrx(ctx)
	err := m.conn.Model(data).Updates(data).Error
	return data, m.maskError(err)
}

func (m *base[T]) DeleteByID(ctx context.Context, id string) error {
	m.checkTrx(ctx)
	err := m.conn.Model(m.entity).Where("id = ?", id).Error
	return m.maskError(err)
}
