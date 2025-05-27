package usecase

import (
	"store/internal/controllers"
	"store/internal/entity"
)

type Usecase struct {
	repo controllers.Repositorio
}

func(u Usecase) Create(cadastro entity.Cadastro) (entity.Cadastro, error) {
	return u.repo.Insert(cadastro)
}

func(u Usecase) Read(cadastro entity.Cadastro) (entity.Cadastro, error) {

	return u.repo.Get(cadastro.Code)
}

func NewUsecase(repo controllers.Repositorio) controllers.Usecase{
	return Usecase{repo: repo}
}	