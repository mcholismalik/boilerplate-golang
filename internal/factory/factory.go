package factory

import (
	"github.com/mcholismalik/boilerplate-golang/internal/factory/repository"
	"github.com/mcholismalik/boilerplate-golang/internal/factory/usecase"
	base "github.com/mcholismalik/boilerplate-golang/internal/model/base"
)

type Factory struct {
	Repository repository.Factory
	Usecase    usecase.Factory
	WsHub      *base.WsHub
}

func Init() Factory {
	f := Factory{}

	f.Repository = repository.Init()
	f.Usecase = usecase.Init(f.Repository)
	f.WsHub = base.NewWsHub()

	return f
}
