package auth

import (
	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	"github.com/mcholismalik/boilerplate-golang/internal/dto"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/internal/middleware"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	Factory factory.Factory
}

func NewHandler(f factory.Factory) *handler {
	return &handler{f}
}

func (h *handler) Route(g *echo.Group) {
	g.POST("/login", h.Login, middleware.Context)
	g.POST("/register", h.Register, middleware.Context)
}

// Login
// @Summary Login user
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.AuthLoginRequest true "request body"
// @Success 200 {object} dto.AuthLoginResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /auth/login [post]
func (h *handler) Login(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(*abstraction.Context)

	payload := new(dto.AuthLoginRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.Factory.Usecase.Auth.Login(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

// Register
// @Summary Register user
// @Description Register user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.AuthRegisterRequest true "request body"
// @Success 200 {object} dto.AuthLoginResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /auth/register [post]
func (h *handler) Register(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(*abstraction.Context)

	payload := new(dto.AuthRegisterRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.Factory.Usecase.Auth.Register(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
