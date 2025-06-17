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

	err := ctx.BodyParser(&product)
	if err != nil {
		log.Printf("Failed to parse product JSON: %v\n", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product JSON",
		})
	}

	log.Printf("Creating product: %+v\n", product)

	result, err := p.product.CreateProduct(product)
	if err != nil {
		log.Printf("Failed to create product: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create product",
		})
	}

	log.Printf("Product created successfully: %+v\n", result)
	return ctx.Status(fiber.StatusOK).JSON(result)

}

func (p productHandler) HandleGetProduct(ctx *fiber.Ctx) error {

	code := ctx.Query("code")

	if code == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "Query parameter 'code' is required"})
	}

	product := entity.Product{Code: code}

	result, err := p.product.GetProduct(product)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(404).JSON(fiber.Map{
				"error": "Product not found",
			})
		}

		log.Printf("Failed to retrieve product: %v\n", err)
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Could not retrieve product",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(result)

}

func (p productHandler) HandleUpdateProduct(ctx *fiber.Ctx) error {

	var product entity.Product

	log.Println("Body recebido:", product)

	err := ctx.BodyParser(&product)
	if err != nil {
		log.Printf("Failed to parse product JSON for update: %v\n", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product JSON",
		})
	}

	log.Printf("Updating product: %+v\n", product)

	result, err := p.product.UpdateProduct(product)
	if err != nil {
		log.Printf("Failed to update product: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update product",
		})
	}

	log.Printf("Product updated successfully: %+v\n", result)
	return ctx.Status(fiber.StatusOK).JSON(result)

}

func (p productHandler) HandleDeleteProduct(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	if code == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product code is required",
		})
	}

	log.Printf("Deleting product with code: %s\n", code)

	if err := p.product.DeleteProduct(code); err != nil {
		log.Printf("Failed to delete product: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete product",
		})
	}

	log.Println("Product deleted successfully")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})

}

func NewProductHandler(product controllers.ProductService) productHandler {
	return productHandler{product: product}
}
