package handlers

import (
	"database/sql"
	"errors"
	"log"
	"store/internal/controllers"
	"store/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	product controllers.ProductService
}

func (p productHandler) HandlePostProduct(ctx *fiber.Ctx) error {

	var product entity.Product

	log.Println("Body recebido:", product)

	err := ctx.BodyParser(&product)
	if err != nil {
		log.Println("Erro ao fazer parse do body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	log.Printf("Produto recebido: %+v\n", product)

	result, err := p.product.CreateProduct(product)
	if err != nil {
		log.Println("Erro ao criar produto:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao adicionar produto",
		})
	}

	log.Println("Produto inserido com sucesso:", result)
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
