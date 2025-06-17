package usecase

import (
	"store/internal/controllers"
	"store/internal/entity"
)

type productService struct {
	repo controllers.Repository
}

func (p productService) CreateProduct(product entity.Product) (entity.Product, error) {

	err := product.ValidateProduct()
	if err != nil {
		return entity.Product{}, err
	}

	return p.repo.Insert(product)
}

func (p productService) GetProduct(product entity.Product) (entity.Product, error) {

	return p.repo.Get(product.Code)
}

func (p productService) UpdateProduct(product entity.Product) (entity.Product, error) {

	return p.repo.Update(product)
}

func (p productService) DeleteProduct(product string) error {

	return p.repo.Delete(product)
}

func NewProductService(repo controllers.Repository) controllers.ProductService {
	return productService{repo: repo}
}
