package handlers

import (
	"store/internal/controllers"
	"store/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	sistem controllers.Sistema
}

func (h Handlers) HandlerPost(ctx *fiber.Ctx) error {
	var consulta entity.Consulta

	err := ctx.BodyParser(&consulta)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Erro ao fazer o parse do corpo", "detalhe": err.Error()})
	}

	if err := consulta.Validator(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resultado, err := h.sistem.CalculoDeItens(consulta)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Erro no cálculo", "detalhe": err.Error()})
	}

	return ctx.Status(200).JSON(resultado)
}

func NewHandlersSistem(sistem controllers.Sistema) Handlers {
	return Handlers{sistem: sistem}
}
