package user

import (
	"context"

	"github.com/mcholismalik/boilerplate-golang/internal/delivery/middleware"
	"github.com/mcholismalik/boilerplate-golang/internal/factory"
	"github.com/mcholismalik/boilerplate-golang/internal/model/abstraction"
	"github.com/mcholismalik/boilerplate-golang/internal/model/dto"
	model "github.com/mcholismalik/boilerplate-golang/internal/model/entity"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
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
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetByID, middleware.Authentication)
	g.POST("", h.Create, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
}

// Get user
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @param request query abstraction.FilterParam true "request query"
// @Param name query string false "name"
// @Success 200 {object} dto.SearchGetResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /rest/users [get]
func (h *delivery) Get(c echo.Context) error {
	filter := abstraction.NewFilterBuiler[model.UserEntity](c, "users")
	if err := c.Bind(filter.Payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(filter.Payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	filter.Bind()

	result, pagination, err := h.Factory.Usecase.User.Find(c.Request().Context(), *filter.Payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result, "Get users success", &pagination).Send(c)
}

// Get user by id
// @Summary Get user by id
// @Description Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /rest/users/{id} [get]
func (h *delivery) GetByID(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(context.Context)

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.Factory.Usecase.User.FindByID(cc, *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Create user
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateUserRequest true "request body"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /rest/users [post]
func (h *delivery) Create(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(context.Context)

	payload := new(dto.CreateUserRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Create(cc, *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update user
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Param request body dto.UpdateUserRequest true "request body"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /rest/users/{id} [put]
func (h *delivery) Update(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(context.Context)

	payload := new(dto.UpdateUserRequest)
	if err := c.Bind(&payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Update(cc, *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete user
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /rest/users/{id} [delete]
func (h *delivery) Delete(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(context.Context)

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Delete(cc, *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
