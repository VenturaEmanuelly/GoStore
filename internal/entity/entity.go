package entity

type Product struct {
	Code  string  `json:"code" validate:"required"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type Order struct {
	Items []Product `json:"items" validate:"required"`
	Total float64   `json:"total"`
}
