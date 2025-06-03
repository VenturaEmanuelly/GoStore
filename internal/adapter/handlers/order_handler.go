package handlers

import (
	"store/internal/controllers"
	"store/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	service controllers.OrderService
}

func (o orderHandler) HandlePostOrder(ctx *fiber.Ctx) error {
	var order entity.Order

	err := ctx.BodyParser(&order)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Error parsing the body", "detail": err.Error()})
	}

	result, err := o.service.CalculateOrder(order)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Error in calculation", "detail": err.Error()})
	}

	return ctx.Status(200).JSON(result)
}

func NewOrderHandler(service controllers.OrderService) orderHandler {
	return orderHandler{service: service}
}
