package usecase

import (
	"github.com/mcholismalik/boilerplate-golang/internal/factory/repository"
	"github.com/mcholismalik/boilerplate-golang/internal/usecase/auth"
	"github.com/mcholismalik/boilerplate-golang/internal/usecase/user"
)

type Factory struct {
	Auth auth.Usecase
	User user.Usecase
}

func Init(r repository.Factory) Factory {
	f := Factory{}
	f.Auth = auth.NewUsecase(r)
	f.User = user.NewUsecase(r)

	return f
}
