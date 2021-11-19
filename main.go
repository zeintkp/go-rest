package main

import (
	"github.com/zeintkp/go-rest/handler"
	"github.com/zeintkp/go-rest/helper/config"
	"github.com/zeintkp/go-rest/helper/database"
	"github.com/zeintkp/go-rest/repository"
	"github.com/zeintkp/go-rest/usecase"

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

	productUsecase := usecase.NewProductUsecase(productRepository)

	handler.NewProductHandler(e, productUsecase)

	e.Logger.Fatal(e.Start(configuration.Get("APP_HOST")))
}
