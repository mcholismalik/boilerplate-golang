package abstraction

import (
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type Filter struct {
	Pagination
	Search string   `query:"search"`
	SortBy []string `query:"sort_by"`
	Query  []FilterQuery
}

type FilterQuery struct {
	Field string
	Value string
}

type FilterBuilder[T any] struct {
	ectx    echo.Context
	name    string
	entity  *T
	Payload *Filter
}

func NewFilterBuiler[T any](ectx echo.Context, name string) FilterBuilder[T] {
	return FilterBuilder[T]{
		ectx:    ectx,
		name:    name,
		entity:  new(T),
		Payload: &Filter{},
	}
}

func (a *FilterBuilder[T]) Bind() {
	req := a.ectx.Request()
	queries := req.URL.Query()
	modelVal := reflect.ValueOf(*a.entity)

	// filter
	ignores := map[string]bool{
		"page":    true,
		"limit":   true,
		"search":  true,
		"sort_by": true,
	}
	for field, values := range queries {
		if ignores[field] {
			continue
		}

		for i := 0; i < modelVal.NumField(); i++ {
			if modelVal.Type().Field(i).Tag.Get("json") == field {
				a.Payload.Query = append(a.Payload.Query, FilterQuery{
					Field: a.name + "." + field,
					Value: values[0],
				})
			}
		}
	}

	// sort
	for _, values := range queries {
		for _, value := range values {
			col := value
			if strings.Contains(value, "-") {
				col = value[:1]
			}

			for i := 0; i < modelVal.NumField(); i++ {
				if modelVal.Type().Field(i).Tag.Get("json") == col {
					a.Payload.SortBy = append(a.Payload.SortBy, value)
				}
			}
		}
	}
}
