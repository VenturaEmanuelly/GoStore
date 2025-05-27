package controllers

import "store/internal/entity"

type RepoSql interface {
	QueryRow(query string, any []any, dest ...any) error
	Exec(query string, args ...any) (int64, error)
}

type Repositorio interface {
	Insert(cadastro entity.Cadastro) (entity.Cadastro, error)
	Get(cadastro string) (entity.Cadastro, error)
	InitSchema() error
}

type Usecase interface {
	Create(cadastro entity.Cadastro) (entity.Cadastro, error)
	Read(cadastro entity.Cadastro) (entity.Cadastro, error)
}

type Sistema interface {
	CalculoDeItens(consulta entity.Consulta) (entity.Consulta, error)
}
