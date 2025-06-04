package entity

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	Code  string  `json:"code" validate:"required"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"gt=0"`
}

type Order struct {
	Items []Product `json:"items" validate:"required"`
	Total float64   `json:"total"`
}

func (o Order) ValidateOrder() error {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		return err
	}

	if len(o.Items) == 0 {
		return fmt.Errorf("it is necessary to inform at least one item")
	}

	return nil
}

func (p Product) ValidateProduct() error {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return err
	}

	return nil
}
