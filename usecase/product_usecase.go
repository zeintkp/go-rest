package usecase

import (
	"database/sql"
	"errors"
	"go-rest/domain"
	pagin "go-rest/helper/pagination"
	"time"

	"github.com/google/uuid"
)

//NewProductUsecase is used to create new Product Usecase
func NewProductUsecase(repository *domain.ProductRepository) domain.ProductUsecase {
	return &ProductUsecase{
		repository: *repository,
	}
}

//ProductUsecase struct
type ProductUsecase struct {
	repository domain.ProductRepository
}

//Browse is used to process browse "products" request
func (uc *ProductUsecase) Browse(search, order, sort string, limit, page int) (productVM []domain.ProductVM, pagination domain.PaginationVM, err error) {
	offset, page, limit, order, sort := pagin.SetPaginationParameter(page, limit, order, sort)

	products, total, err := uc.repository.Browse(search, order, sort, limit, offset)
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
func (uc *ProductUsecase) Create(req domain.ProductRequest) (productVM domain.ProductVM, err error) {

	isExist, err := uc.repository.IsExist("", req.ProductName)
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

	err = uc.repository.Create(product)

	return uc.ConvertVM(product), err
}

//Read is used to process read "products" request
func (uc *ProductUsecase) Read(id string) (productVM domain.ProductVM, err error) {

	product, err := uc.repository.Read(id)
	if err != nil {
		return
	}

	return uc.ConvertVM(product), err
}

//Update is used to process update "products" request
func (uc *ProductUsecase) Update(req domain.ProductRequest, id string) (productVM domain.ProductVM, err error) {

	product, err := uc.repository.Read(id)
	if err != nil {
		return
	}

	isExist, err := uc.repository.IsExist(id, req.ProductName)
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

	err = uc.repository.Update(product)

	return uc.ConvertVM(product), err
}

//Delete is used to process delete "products" request
func (uc *ProductUsecase) Delete(id string) (productVM domain.ProductVM, err error) {

	product, err := uc.repository.Read(id)
	if err != nil {
		return
	}

	product.DeletedAt = sql.NullString{
		String: time.Now().UTC().Format(time.RFC3339),
		Valid:  true,
	}

	err = uc.repository.Delete(product)

	return uc.ConvertVM(product), err
}

//ConvertVM
func (uc *ProductUsecase) ConvertVM(product domain.Product) domain.ProductVM {
	return domain.ProductVM{
		ID:          product.ID,
		ProductName: product.ProductName,
		IsActive:    product.IsActive,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt.String,
		DeletedAt:   product.DeletedAt.String,
	}
}
