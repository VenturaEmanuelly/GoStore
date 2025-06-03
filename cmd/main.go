package main

import (
	driver "database/sql"
	"fmt"
	"log"
	"os"
	"store/internal/adapter/handlers"
	"store/internal/adapter/repository"
	"store/internal/controllers"
	usecase "store/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	product, service, err := injectDependency()

	if err != nil {
		log.Fatal(err)
	}

	handlerProduct := handlers.NewProductHandler(product)

	orderService := usecase.NewOrderServices(service)
	handlerService := handlers.NewOrderHandler(orderService)

	app := fiber.New()
	app.Post("/product", handlerProduct.HandlePostProduct)
	app.Post("/service", handlerService.HandlePostOrder)
	app.Get("/product", handlerProduct.HandleGetProduct)

	log.Fatal(app.Listen(":8080"))
}

func injectDependency() (controllers.ProductService, controllers.Repository, error) {

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

	repoSql := repository.NewDatabase(db)
	repository := repository.Newrepository(repoSql)

	if err := repository.InitSchema(); err != nil {
		return nil, nil, err
	}

	product := usecase.NewUsecase(repository)

	return product, repository, nil
}
