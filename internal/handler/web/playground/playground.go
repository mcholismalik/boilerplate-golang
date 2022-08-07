package playground

import (
	"net/http"

	"github.com/google/uuid"
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
	key := uuid.New()
	type M map[string]interface{}
	data := M{"message": "playground", "key": key}
	return c.Render(http.StatusOK, "playground.html", data)
}
