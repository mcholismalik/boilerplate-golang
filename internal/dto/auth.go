package dto

import (
	"github.com/mcholismalik/boilerplate-golang/internal/model"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"
)

// request
type (
	AuthLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	AuthRegisterRequest struct {
		model.UserEntity
	}
)

// response
type (
	AuthLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
		model.UserEntityModel
	}
	AuthLoginResponseDoc struct {
		Body struct {
			Meta res.Meta          `json:"meta"`
			Data AuthLoginResponse `json:"data"`
		} `json:"body"`
	}

	AuthRegisterResponse struct {
		model.UserEntityModel
	}
	AuthRegisterResponseDoc struct {
		Body struct {
			Meta res.Meta             `json:"meta"`
			Data AuthRegisterResponse `json:"data"`
		} `json:"body"`
	}
)
