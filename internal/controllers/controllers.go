package controllers

import "store/internal/entity"

type RepoSql interface {
	QueryRow(query string, args []any, dest ...any) error
	Exec(query string, args ...any) (int64, error)
}

type Repository interface {
	Insert(product entity.Product) (entity.Product, error)
	Get(product string) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(product string) error
	InitSchema() error
}

type ProductService interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	GetProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(product string) error
}

type OrderService interface {
	CalculateOrder(order entity.Order) (entity.Order, error)
}
