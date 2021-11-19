package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/zeintkp/go-rest/domain"
	pagin "github.com/zeintkp/go-rest/helper/pagination"

	"github.com/google/uuid"
)

//NewProductUsecase is used to create new Product Usecase
func NewProductUsecase(repository domain.ProductRepository) *productUsecase {
	return &productUsecase{
		repository: repository,
	}
}

//productUsecase struct
type productUsecase struct {
	repository domain.ProductRepository
}

//Browse is used to process browse "products" request
func (uc *productUsecase) Browse(ctx context.Context, search, order, sort string, limit, page int) (productVM []domain.ProductVM, pagination domain.PaginationVM, err error) {
	offset, page, limit, order, sort := pagin.SetPaginationParameter(page, limit, order, sort)

	products, total, err := uc.repository.Browse(ctx, search, order, sort, limit, offset)
	if err != nil {
		return
	}

	if total > 0 {
		for _, data := range products {
			productVM = append(productVM, uc.ConvertVM(data))
		}
	}

	pagination = pagin.SetPaginationResponse(page, limit, total)

	return
}

//Create is used to process create "products" request
func (uc *productUsecase) Create(ctx context.Context, req domain.ProductRequest) (productVM domain.ProductVM, err error) {

	isExist, err := uc.repository.IsExist(ctx, "", req.ProductName)
	if err != nil {
		return
	}

	if isExist {
		return productVM, errors.New(req.ProductName + " already exist")
	}

	product := domain.Product{
		ID:          uuid.New().String(),
		ProductName: req.ProductName,
		IsActive:    true,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
	}

	err = uc.repository.Create(ctx, product)

	return uc.ConvertVM(product), err
}

//Read is used to process read "products" request
func (uc *productUsecase) Read(ctx context.Context, id string) (productVM domain.ProductVM, err error) {

	product, err := uc.repository.Read(ctx, id)
	if err != nil {
		return
	}

	return uc.ConvertVM(product), err
}

//Update is used to process update "products" request
func (uc *productUsecase) Update(ctx context.Context, req domain.ProductRequest, id string) (productVM domain.ProductVM, err error) {

	product, err := uc.repository.Read(ctx, id)
	if err != nil {
		return
	}

	isExist, err := uc.repository.IsExist(ctx, id, req.ProductName)
	if err != nil {
		return
	}

	if isExist {
		return productVM, errors.New(req.ProductName + " already exist")
	}

	product.ProductName = req.ProductName
	product.IsActive = product.IsActive
	product.UpdatedAt = sql.NullString{
		String: time.Now().UTC().Format(time.RFC3339),
		Valid:  true,
	}

	err = uc.repository.Update(ctx, product)

	return uc.ConvertVM(product), err
}

//Delete is used to process delete "products" request
func (uc *productUsecase) Delete(ctx context.Context, id string) (productVM domain.ProductVM, err error) {

	product, err := uc.repository.Read(ctx, id)
	if err != nil {
		return
	}

	product.DeletedAt = sql.NullString{
		String: time.Now().UTC().Format(time.RFC3339),
		Valid:  true,
	}

	err = uc.repository.Delete(ctx, product)

	return uc.ConvertVM(product), err
}

//ConvertVM
func (uc *productUsecase) ConvertVM(product domain.Product) domain.ProductVM {
	return domain.ProductVM{
		ID:          product.ID,
		ProductName: product.ProductName,
		IsActive:    product.IsActive,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt.String,
		DeletedAt:   product.DeletedAt.String,
	}
}
