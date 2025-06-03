package usecase

import (
	"store/internal/controllers"
	"store/internal/entity"

	"github.com/go-playground/validator/v10"
)

type productService struct {
	repo controllers.Repository
}

func (p productService) CreateProduct(product entity.Product) (entity.Product, error) {

	err := p.ValidateProduct(product)
	if err != nil {
		return entity.Product{}, err
	}

	return p.repo.Insert(product)
}

func (p productService) GetProduct(product entity.Product) (entity.Product, error) {

	return p.repo.Get(product.Code)
}

func (p productService) ValidateProduct(product entity.Product) error {
	validate := validator.New()

	err := validate.Struct(product)
	if err != nil {
		return err
	}

	return nil
}

func NewUsecase(repo controllers.Repository) controllers.ProductService {
	return productService{repo: repo}
}
