package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zeintkp/go-rest/domain"
	"github.com/zeintkp/go-rest/domain/mocks"
)

func TestBrowse(t *testing.T) {

	var mockProductVM domain.ProductVM
	err := faker.FakeData(&mockProductVM)
	assert.NoError(t, err)

	var mockPaginVM domain.PaginationVM
	err = faker.FakeData(&mockPaginVM)
	assert.NoError(t, err)

	mockListProductVM := make([]domain.ProductVM, 0)
	mockListProductVM = append(mockListProductVM, mockProductVM)

	search := "pro"
	order := "id"
	sort := "asc"
	limit := 10
	page := 1

	mockUC := new(mocks.ProductUsecase)
	mockUC.On("Browse", mock.Anything, search, order, sort, limit, page).Return(mockListProductVM, mockPaginVM, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("/api/v1/products?search=%s&order=%s&sort=%s&limit=%d&page=%d", search, order, sort, limit, page), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := ProductHandler{
		productUc: mockUC,
	}

	err = handler.Browse(c)
	require.NoError(t, err)
	assert.NotEmpty(t, rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestCreate(t *testing.T) {

	mockProductRequest := domain.ProductRequest{
		ProductName: "test",
		IsActive:    true,
	}

	mockProductVM := domain.ProductVM{
		ID:          uuid.New().String(),
		ProductName: mockProductRequest.ProductName,
		IsActive:    mockProductRequest.IsActive,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	j, err := json.Marshal(mockProductRequest)
	assert.NoError(t, err)

	mockUC := new(mocks.ProductUsecase)
	mockUC.On("Create", mock.Anything, mockProductRequest).Return(mockProductVM, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/api/v1/products", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := ProductHandler{
		productUc: mockUC,
	}

	err = handler.Create(c)
	require.NoError(t, err)
	assert.NotEmpty(t, rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestRead(t *testing.T) {
	var mockProductVM domain.ProductVM
	err := faker.FakeData(&mockProductVM)
	assert.NoError(t, err)

	id := uuid.New().String()

	mockUC := new(mocks.ProductUsecase)
	mockUC.On("Read", mock.Anything, id).Return(mockProductVM, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/api/v1/products/"+id, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/products/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)
	handler := ProductHandler{
		productUc: mockUC,
	}

	err = handler.Read(c)
	require.NoError(t, err)
	assert.NotEmpty(t, rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {

	id := uuid.New().String()

	mockProductRequest := domain.ProductRequest{
		ProductName: "test",
		IsActive:    true,
	}

	mockProductVM := domain.ProductVM{
		ID:          id,
		ProductName: mockProductRequest.ProductName,
		IsActive:    mockProductRequest.IsActive,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	j, err := json.Marshal(mockProductRequest)
	assert.NoError(t, err)

	mockUC := new(mocks.ProductUsecase)
	mockUC.On("Update", mock.Anything, mockProductRequest, id).Return(mockProductVM, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/api/v1/products/"+id, strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/products/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)
	handler := ProductHandler{
		productUc: mockUC,
	}

	err = handler.Update(c)
	require.NoError(t, err)
	assert.NotEmpty(t, rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestDelete(t *testing.T) {

	id := uuid.New().String()

	mockProductVM := domain.ProductVM{
		ID:          id,
		ProductName: "test",
		IsActive:    true,
		CreatedAt:   time.Now().Format(time.RFC3339),
		DeletedAt:   time.Now().Format(time.RFC3339),
	}

	mockUC := new(mocks.ProductUsecase)
	mockUC.On("Delete", mock.Anything, id).Return(mockProductVM, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/api/v1/products/"+id, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/products/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)
	handler := ProductHandler{
		productUc: mockUC,
	}

	err = handler.Delete(c)
	require.NoError(t, err)
	assert.NotEmpty(t, rec.Body)
	fmt.Println(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}
