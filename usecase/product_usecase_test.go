package usecase

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zeintkp/go-rest/domain"
	"github.com/zeintkp/go-rest/domain/mocks"
)

func TestBrowse(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{
		ID:          uuid.New().String(),
		ProductName: "test",
		IsActive:    true,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	mockListProduct := make([]domain.Product, 0)
	mockListProduct = append(mockListProduct, mockProduct)

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Browse", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockListProduct, 1, nil).Once()

		search := "test"
		order := "id"
		sort := "asc"
		limit := 10
		page := 1

		uc := NewProductUsecase(mockProductRepo)
		list, pagin, err := uc.Browse(context.TODO(), search, order, sort, limit, page)
		assert.NoError(t, err)
		assert.Equal(t, 1, pagin.Total)
		assert.Len(t, list, len(mockListProduct))

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockProductRepo.On("Browse", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, 0, errors.New("Unexpected Error")).Once()

		search := "test"
		order := "id"
		sort := "asc"
		limit := 10
		page := 1

		uc := NewProductUsecase(mockProductRepo)
		list, pagin, err := uc.Browse(context.TODO(), search, order, sort, limit, page)
		assert.Error(t, err)
		assert.Empty(t, pagin)
		assert.Len(t, list, 0)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRequest := domain.ProductRequest{
		ProductName: "test",
		IsActive:    true,
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("IsExist", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(false, nil).Once()

		mockProductRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Product")).
			Return(nil).Once()

		uc := NewProductUsecase(mockProductRepo)

		productVM, err := uc.Create(context.TODO(), mockProductRequest)

		assert.NoError(t, err)
		assert.Equal(t, productVM.ProductName, mockProductRequest.ProductName)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("existing-product", func(t *testing.T) {
		mockProductRepo.On("IsExist", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(true, nil).Once()

		uc := NewProductUsecase(mockProductRepo)

		productVM, err := uc.Create(context.TODO(), mockProductRequest)

		assert.Error(t, err)
		assert.Empty(t, productVM.ProductName)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestRead(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)

	id := uuid.New().String()
	mockProduct := domain.Product{
		ID:          id,
		ProductName: "test",
		IsActive:    true,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Read", mock.Anything, mock.AnythingOfType("string")).
			Return(mockProduct, nil).Once()

		uc := NewProductUsecase(mockProductRepo)
		data, err := uc.Read(context.TODO(), id)

		assert.NoError(t, err)
		assert.Equal(t, data.ID, id)

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockProductRepo.On("Read", mock.Anything, mock.AnythingOfType("string")).
			Return(domain.Product{}, sql.ErrNoRows).Once()

		uc := NewProductUsecase(mockProductRepo)
		data, err := uc.Read(context.TODO(), id)

		assert.Error(t, err)
		assert.Empty(t, data.ID)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRequest := domain.ProductRequest{
		ProductName: "test2",
		IsActive:    true,
	}
	id := uuid.New().String()
	mockProduct := domain.Product{
		ID:          id,
		ProductName: "test",
		IsActive:    true,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Read", mock.Anything, mock.AnythingOfType("string")).
			Return(mockProduct, nil).Once()

		mockProductRepo.On("IsExist", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(false, nil).Once()

		mockProductRepo.On("Update", mock.Anything, mock.AnythingOfType("domain.Product")).
			Return(nil).Once()

		uc := NewProductUsecase(mockProductRepo)
		productVM, err := uc.Update(context.TODO(), mockProductRequest, id)
		assert.NoError(t, err)
		assert.NotEmpty(t, productVM.UpdatedAt)
		assert.Equal(t, mockProductRequest.ProductName, productVM.ProductName)

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("existing-name", func(t *testing.T) {
		mockProductRepo.On("Read", mock.Anything, mock.AnythingOfType("string")).
			Return(mockProduct, nil).Once()

		mockProductRepo.On("IsExist", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(true, nil).Once()

		uc := NewProductUsecase(mockProductRepo)
		productVM, err := uc.Update(context.TODO(), mockProductRequest, id)
		assert.Error(t, err)
		assert.Empty(t, productVM.UpdatedAt)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)

	id := uuid.New().String()
	mockProduct := domain.Product{
		ID:          id,
		ProductName: "test",
		IsActive:    true,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Read", mock.Anything, mock.AnythingOfType("string")).
			Return(mockProduct, nil).Once()

		mockProductRepo.On("Delete", mock.Anything, mock.AnythingOfType("domain.Product")).
			Return(nil).Once()

		uc := NewProductUsecase(mockProductRepo)
		productVM, err := uc.Delete(context.TODO(), id)
		assert.NoError(t, err)
		assert.NotEmpty(t, productVM.DeletedAt)

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("id-not-found", func(t *testing.T) {
		mockProductRepo.On("Read", mock.Anything, mock.AnythingOfType("string")).
			Return(domain.Product{}, sql.ErrNoRows).Once()

		uc := NewProductUsecase(mockProductRepo)
		productVM, err := uc.Delete(context.TODO(), id)
		assert.Error(t, err)
		assert.Empty(t, productVM.DeletedAt)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestConvertVM(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{
		ID:          uuid.New().String(),
		ProductName: "test",
		IsActive:    true,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	uc := NewProductUsecase(mockProductRepo)
	productVM := uc.ConvertVM(mockProduct)
	assert.Equal(t, mockProduct.ID, productVM.ID)
	assert.Equal(t, mockProduct.ProductName, productVM.ProductName)
	assert.Equal(t, mockProduct.IsActive, productVM.IsActive)
	assert.Equal(t, mockProduct.CreatedAt, productVM.CreatedAt)

	mockProductRepo.AssertExpectations(t)
}

func Test_productUsecase_Browse(t *testing.T) {
	type fields struct {
		repository domain.ProductRepository
	}
	type args struct {
		ctx    context.Context
		search string
		order  string
		sort   string
		limit  int
		page   int
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantProductVM  []domain.ProductVM
		wantPagination domain.PaginationVM
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &productUsecase{
				repository: tt.fields.repository,
			}
			gotProductVM, gotPagination, err := uc.Browse(tt.args.ctx, tt.args.search, tt.args.order, tt.args.sort, tt.args.limit, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("productUsecase.Browse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProductVM, tt.wantProductVM) {
				t.Errorf("productUsecase.Browse() gotProductVM = %v, want %v", gotProductVM, tt.wantProductVM)
			}
			if !reflect.DeepEqual(gotPagination, tt.wantPagination) {
				t.Errorf("productUsecase.Browse() gotPagination = %v, want %v", gotPagination, tt.wantPagination)
			}
		})
	}
}
