package handlers

import (
	"database/sql"
	"errors"
	"store/internal/controllers"
	"store/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	product controllers.ProductService
}

func (p productHandler) HandlePostProduct(ctx *fiber.Ctx) error {
	var product entity.Product
	err := ctx.BodyParser(&product)
	if err != nil {
		return err
	}

	result, err := p.product.CreateProduct(product)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(result)
}

func (p productHandler) HandleGetProduct(ctx *fiber.Ctx) error {

	code := ctx.Query("code")

	if code == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "Parameter 'code' is mandatory"})
	}

	product := entity.Product{Code: code}

	result, err := p.product.GetProduct(product)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(404).JSON(fiber.Map{"error": "registration not found"})
		}
		return ctx.Status(500).JSON(fiber.Map{"error": "error retrieving registration"})
	}
	return ctx.Status(200).JSON(result)

}

func NewProductHandler(product controllers.ProductService) productHandler {
	return productHandler{product: product}
}
