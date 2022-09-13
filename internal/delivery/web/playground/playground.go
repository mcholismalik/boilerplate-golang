package playground

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	Factory factory.Factory
}

func NewDelivery(f factory.Factory) *delivery {
	return &delivery{f}
}

func (h *delivery) Route(g *echo.Group) {
	g.GET("", h.Get)
}

func (h *delivery) Get(c echo.Context) error {
	key := uuid.New()
	type M map[string]interface{}
	data := M{"message": "playground", "key": key}
	return c.Render(http.StatusOK, "playground.html", data)
}
