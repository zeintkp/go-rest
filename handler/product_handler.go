package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/zeintkp/go-rest/domain"
	"github.com/zeintkp/go-rest/helper/requestHelper"
	"github.com/zeintkp/go-rest/helper/responseHelper"

	"github.com/labstack/echo/v4"
)

//NewProductHandler is used to create Product Handler
func NewProductHandler(e *echo.Echo, productUC domain.ProductUsecase) {
	handler := &ProductHandler{
		productUc: productUC,
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
func (handler *ProductHandler) Browse(c echo.Context) error {

	search := c.QueryParam("search")
	order := c.QueryParam("order")
	sort := c.QueryParam("sort")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ctx := c.Request().Context()

	productVM, paginationVM, err := handler.productUc.Browse(ctx, search, order, sort, limit, page)

	return c.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, paginationVM, err))
}

//Create is used to handle create in products endpoint
func (handler *ProductHandler) Create(c echo.Context) error {

	input, err := requestHelper.GetInstance().ValidateRequest(c, new(domain.ProductRequest))
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(responseHelper.BuildResponse(http.StatusBadRequest, nil, nil, err))
	}
	ctx := c.Request().Context()

	productVM, err := handler.productUc.Create(ctx, *input.(*domain.ProductRequest))
	return c.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, nil, err))
}

//Read is used to handle read in products endpoint
func (handler *ProductHandler) Read(c echo.Context) error {

	id := c.Param("id")
	ctx := c.Request().Context()

	productVM, err := handler.productUc.Read(ctx, id)
	return c.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, nil, err))
}

//Update is used to handle update in products endpoint
func (handler *ProductHandler) Update(c echo.Context) error {

	input, err := requestHelper.GetInstance().ValidateRequest(c, new(domain.ProductRequest))
	if err != nil {
		return c.JSON(responseHelper.BuildResponse(http.StatusBadRequest, nil, nil, err))
	}

	id := c.Param("id")
	ctx := c.Request().Context()

	productVM, err := handler.productUc.Update(ctx, *input.(*domain.ProductRequest), id)
	return c.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, nil, err))
}

//Delete is used to handle delete in products endpoint
func (handler *ProductHandler) Delete(c echo.Context) error {

	id := c.Param("id")
	ctx := c.Request().Context()

	productVM, err := handler.productUc.Delete(ctx, id)
	return c.JSON(responseHelper.BuildResponse(http.StatusOK, productVM, nil, err))
}
