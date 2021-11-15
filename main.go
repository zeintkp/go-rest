package main

import (
	"go-rest/handler"
	"go-rest/helper/config"
	"go-rest/helper/database"
	"go-rest/repository"
	"go-rest/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	//Setup configuration
	configuration := config.New("./.env")

	//Create echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Setup DB connection
	db := database.NewPostgresDatabase(configuration)

	productRepository := repository.NewProductRepository(db)

	productUsecase := usecase.NewProductUsecase(&productRepository)

	handler.NewProductHandler(e, &productUsecase)

	e.Logger.Fatal(e.Start(configuration.Get("APP_HOST")))
}
