package auth

import (
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/internal/model/dto"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	Factory factory.Factory
}

func NewDelivery(f factory.Factory) *delivery {
	return &delivery{f}
}

func (h *delivery) Route(g *echo.Group) {
	g.POST("/login", h.Login)
	g.POST("/register", h.Register)
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
// @Router /rest/auth/login [post]
func (h *delivery) Login(c echo.Context) error {
	payload := new(dto.AuthLoginRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.Factory.Usecase.Auth.Login(c.Request().Context(), *payload)
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
// @Router /rest/auth/register [post]
func (h *delivery) Register(c echo.Context) error {
	payload := new(dto.AuthRegisterRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.Factory.Usecase.Auth.Register(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
