package abstraction

import (
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type FilterParam struct {
	Pagination
	Search string   `query:"search"`
	SortBy []string `query:"sort_by"`
	Query  []QueryFilter
}

type QueryFilter struct {
	Field string
	Value string
}

func BindFilterSort[T any](c echo.Context, model T, name string, payload *FilterParam) {
	BindFilter(c, model, name, payload)
	BindSort(c, model, name, payload)
}

func BindFilter[T any](c echo.Context, model T, name string, payload *FilterParam) {
	var filters []QueryFilter

	req := c.Request()
	queries := req.URL.Query()
	modelVal := reflect.ValueOf(model)

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
				filters = append(filters, QueryFilter{
					Field: name + "." + field,
					Value: values[0],
				})
			}
		}
	}

	payload.Query = filters
}

func BindSort[T any](c echo.Context, model T, name string, payload *FilterParam) {
	var sortBy []string

	req := c.Request()
	queries := req.URL.Query()
	modelVal := reflect.ValueOf(model)

	for _, values := range queries {
		for _, value := range values {
			col := value
			if strings.Contains(value, "-") {
				col = value[:1]
			}

			for i := 0; i < modelVal.NumField(); i++ {
				if modelVal.Type().Field(i).Tag.Get("json") == col {
					sortBy = append(sortBy, value)
				}
			}
		}
	}

	payload.SortBy = sortBy
}
