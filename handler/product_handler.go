package handler

import (
	"fmt"
	"go-rest/domain"
	"go-rest/helper/requestHelper"
	"go-rest/helper/responseHelper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//NewProductHandler is used to create Product Handler
func NewProductHandler(e *echo.Echo, productUC *domain.ProductUsecase) {
	handler := &ProductHandler{
		productUc: *productUC,
	}

	routeV1(e.Group("/api/v1/products"), handler)
}

//RegisterProductRoute is used to register endpoints to access products
func routeV1(route *echo.Group, handler *ProductHandler) {
	route.GET("", handler.Browse)
	route.POST("", handler.Create)
	route.GET("/:id", handler.Read)
	route.PUT("/:id", handler.Update)
	route.DELETE("/:id", handler.Delete)
}

//ProductHandler struct
type ProductHandler struct {
	productUc domain.ProductUsecase
}

//Browse is used to handle browse in products endpoint
func (handler *ProductHandler) Browse(ctx echo.Context) error {

	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	productVM, paginationVM, err := handler.productUc.Browse(search, order, sort, limit, page)

	return ctx.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, paginationVM, err))
}

//Create is used to handle create in products endpoint
func (handler *ProductHandler) Create(ctx echo.Context) error {

	input, err := requestHelper.GetInstance().ValidateRequest(ctx, new(domain.ProductRequest))
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(responseHelper.BuildResponse(http.StatusBadRequest, nil, nil, err))
	}

	productVM, err := handler.productUc.Create(*input.(*domain.ProductRequest))
	return ctx.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, nil, err))
}

//Read is used to handle read in products endpoint
func (handler *ProductHandler) Read(ctx echo.Context) error {

	id := ctx.Param("id")

	productVM, err := handler.productUc.Read(id)
	return ctx.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, nil, err))
}

//Update is used to handle update in products endpoint
func (handler *ProductHandler) Update(ctx echo.Context) error {

	input, err := requestHelper.GetInstance().ValidateRequest(ctx, new(domain.ProductRequest))
	if err != nil {
		return ctx.JSON(responseHelper.BuildResponse(http.StatusBadRequest, nil, nil, err))
	}

	id := ctx.Param("id")

	productVM, err := handler.productUc.Update(*input.(*domain.ProductRequest), id)
	return ctx.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, nil, err))
}

//Delete is used to handle delete in products endpoint
func (handler *ProductHandler) Delete(ctx echo.Context) error {

	id := ctx.Param("id")

	productVM, err := handler.productUc.Delete(id)
	return ctx.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, nil, err))
}
