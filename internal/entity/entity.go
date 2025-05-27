package entity

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Cadastro struct {
	Code  string  `json:"code" validate:"required"`
	Iten  string  `json:"iten" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type Consulta struct {
	Iten  []Cadastro `json:"iten" validate:"required"`
	Total float64    `json:"total"`
}

func (c Cadastro) Validator() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

func (c Consulta) Validator() error {
	validate := validator.New()
	
	err := validate.Struct(c)
	if err != nil {
		return err
	}

	
	if len(c.Iten) == 0 {
		return fmt.Errorf("é necessário informar pelo menos um item")
	}

	return nil
}
