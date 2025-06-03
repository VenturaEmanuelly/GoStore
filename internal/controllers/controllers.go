package controllers

import "store/internal/entity"

type RepoSql interface {
	QueryRow(query string, any []any, dest ...any) error
	Exec(query string, args ...any) (int64, error)
}

type Repository interface {
	Insert(product entity.Product) (entity.Product, error)
	Get(product string) (entity.Product, error)
	InitSchema() error
}

type ProductService interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	GetProduct(product entity.Product) (entity.Product, error)
}

type OrderService interface {
	CalculateOrder(order entity.Order) (entity.Order, error)
}
