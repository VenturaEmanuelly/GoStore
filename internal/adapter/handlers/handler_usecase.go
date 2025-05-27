package handlers

import (
	"database/sql"
	"errors"
	"store/internal/controllers"
	"store/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase controllers.Usecase
}

func (h Handler) HandlerPost(ctx *fiber.Ctx) error {
	var cadastro entity.Cadastro
	err := ctx.BodyParser(&cadastro)
	if err != nil {
		return err
	}

	if err := cadastro.Validator(); err != nil {
    return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error": err.Error(),
    })
}

	resultado, err := h.usecase.Create(cadastro)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(resultado)
}

func (h Handler) HandlerGet(ctx *fiber.Ctx) error {

	code := ctx.Query("code")

	if code == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "Parâmetro 'code' é obrigatório"})
	}

	cadastro := entity.Cadastro{Code: code}

	resultado, err := h.usecase.Read(cadastro)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.Status(404).JSON(fiber.Map{"error": "Cadastro não encontrado"})
		}
		return ctx.Status(500).JSON(fiber.Map{"error": "Erro ao buscar cadastro"})
	}
	return ctx.Status(200).JSON(resultado)

}

func NewHandlerUsecase(usecase controllers.Usecase) Handler {
	return Handler{usecase: usecase}
}
