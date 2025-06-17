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

	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Invalid request body",
			"detail": err.Error(),
		})
	}

	result, err := o.service.CalculateOrder(order)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to calculate order",
			"detail": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}


func NewOrderHandler(service controllers.OrderService) orderHandler {
	return orderHandler{service: service}
}
