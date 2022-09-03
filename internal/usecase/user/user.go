package user

import (
	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	"github.com/mcholismalik/boilerplate-golang/internal/dto"
	"github.com/mcholismalik/boilerplate-golang/internal/factory/repository"
	"github.com/mcholismalik/boilerplate-golang/internal/model"
	res "github.com/mcholismalik/boilerplate-golang/pkg/util/response"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/str"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/trxmanager"

	"github.com/google/uuid"
)

type Usecase interface {
	Find(ctx abstraction.Context, filterParam *abstraction.FilterParam) ([]dto.UserResponse, abstraction.PaginationInfo, error)
	FindByID(ctx abstraction.Context, payload dto.ByIDRequest) (dto.UserResponse, error)
	Create(ctx abstraction.Context, payload dto.CreateUserRequest) (dto.UserResponse, error)
	Update(ctx abstraction.Context, payload dto.UpdateUserRequest) (dto.UserResponse, error)
	Delete(ctx abstraction.Context, payload dto.ByIDRequest) (dto.UserResponse, error)
}

type usecase struct {
	RepositoryFactory repository.Factory
}

func NewUsecase(f repository.Factory) *usecase {
	return &usecase{f}
}

func (u *usecase) Find(ctx abstraction.Context, filterParam *abstraction.FilterParam) (result []dto.UserResponse, pagination abstraction.PaginationInfo, err error) {
	users, info, err := u.RepositoryFactory.UserRepository.Find(ctx, filterParam)
	if err != nil {
		return nil, pagination, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	pagination = *info

	for _, user := range users {
		result = append(result, dto.UserResponse{
			Data: user,
		})
	}

	return result, pagination, nil
}

func (u *usecase) FindByID(ctx abstraction.Context, payload dto.ByIDRequest) (dto.UserResponse, error) {
	var result dto.UserResponse

	data, err := u.RepositoryFactory.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {
		return result, err
	}

	result = dto.UserResponse{
		Data: *data,
	}

	return result, nil
}

func (u *usecase) Create(ctx abstraction.Context, payload dto.CreateUserRequest) (result dto.UserResponse, err error) {
	var email string
	if payload.Email != nil {
		email = *payload.Email
	} else {
		email = str.GenerateRandString(10) + "@gmail.com"
	}

	var (
		data   *model.UserModel
		userID = uuid.New()
		user   = model.UserModel{
			Entity: abstraction.Entity{
				ID: userID,
			},
			UserEntity: model.UserEntity{
				Name:     payload.Name,
				Email:    email,
				Password: payload.Password,
			},
			Context: ctx,
		}
	)

	if err = trxmanager.New(u.RepositoryFactory.Db).WithTrx(ctx, func(ctx abstraction.Context) error {
		data, err = u.RepositoryFactory.UserRepository.Create(ctx, &user)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.UserResponse{
		Data: *data,
	}

	return result, nil
}

func (u *usecase) Update(ctx abstraction.Context, payload dto.UpdateUserRequest) (result dto.UserResponse, err error) {
	var (
		data *model.UserModel
		user = model.UserModel{
			UserEntity: model.UserEntity{
				Name:     payload.Name,
				Email:    payload.Email,
				Password: payload.Password,
			},
			Context: ctx,
		}
	)

	if payload.Password != "" {
		user.HashPassword()
		user.Password = ""
	}

	if err = trxmanager.New(u.RepositoryFactory.Db).WithTrx(ctx, func(ctx abstraction.Context) error {
		data, err = u.RepositoryFactory.UserRepository.UpdateByID(ctx, payload.ID, &user)
		if err != nil {
			return err
		}

		data, err = u.RepositoryFactory.UserRepository.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.UserResponse{
		Data: *data,
	}

	return result, nil
}

func (u *usecase) Delete(ctx abstraction.Context, payload dto.ByIDRequest) (result dto.UserResponse, err error) {
	var data *model.UserModel

	if err = trxmanager.New(u.RepositoryFactory.Db).WithTrx(ctx, func(ctx abstraction.Context) error {
		data, err = u.RepositoryFactory.UserRepository.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		err = u.RepositoryFactory.UserRepository.DeleteByID(ctx, payload.ID)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = dto.UserResponse{
		Data: *data,
	}

	return result, nil
}