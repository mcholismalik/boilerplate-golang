package bubble

import (
	"net/http"

	"github.com/mcholismalik/boilerplate-golang/internal/factory"

	"github.com/labstack/echo/v4"
)

type handler struct {
	Factory factory.Factory
}

func NewHandler(f factory.Factory) *handler {
	return &handler{f}
}

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
}

func (h *handler) Get(c echo.Context) error {
	key := ""
	queries := c.Request().URL.Query()
	for field, values := range queries {
		if field == "key" {
			if len(values) == 0 {
				continue
			}
			key = values[0]
		}
	}

	type M map[string]interface{}
	data := M{"message": "bubble", "key": key}
	return c.Render(http.StatusOK, "bubble.html", data)
}
