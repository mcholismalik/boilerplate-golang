package dto

import (
	model "github.com/mcholismalik/boilerplate-golang/internal/model/entity"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"
)

// request
type (
	CreateUserRequest struct {
		Name     string  `json:"name" validate:"required"`
		Email    *string `json:"email,omitempty" validate:"omitempty"`
		Password string  `json:"password"`
	}
)

type (
	UpdateUserRequest struct {
		ID       string `param:"id" validate:"required"`
		Name     string `json:"name"`
		Email    string `json:"email" validate:"omitempty,email"`
		Password string `json:"password"`
	}
)

// response
type (
	UserResponse struct {
		Data model.UserModel
	}
	UserResponseDoc struct {
		Body struct {
			Meta res.Meta     `json:"meta"`
			Data UserResponse `json:"data"`
		} `json:"body"`
	}
)
