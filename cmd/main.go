package main

import (
	driver "database/sql"
	"fmt"
	"log"
	"os"
	"store/internal/adapter/handlers"
	"store/internal/adapter/repositorio"
	"store/internal/controllers"
	usecase "store/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	usecaseCadastro,repoCadastro, err := injectDependency()

	if err != nil {
		log.Fatal(err)
	}

	handlerCadastro := handlers.NewHandlerUsecase(usecaseCadastro)

	consulta := usecase.NewSistemas(repoCadastro)
	handlerConsulta := handlers.NewHandlersSistem(consulta)

	app := fiber.New()
	app.Post("/cadastro", handlerCadastro.HandlerPost)
	app.Post("/consulta", handlerConsulta.HandlerPost)
	app.Get("/cadastro", handlerCadastro.HandlerGet)

	log.Fatal(app.Listen(":8080"))
}

func injectDependency() (controllers.Usecase,controllers.Repositorio, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	db, err := driver.Open("postgres", psqlInfo)
	if err != nil {
		return nil, nil, err
	}

	repo := repositorio.NewDatabase(db)
	cadastro := repositorio.NewRepositorio(repo)

	if err := cadastro.InitSchema(); err != nil {
		return nil, nil, err
	}

	usecaseCadastro := usecase.NewUsecase(cadastro)

	return usecaseCadastro, cadastro, nil
}
