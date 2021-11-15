package domain

import "database/sql"

//Products struct
type Product struct {
	ID          string         `db:"id"`
	ProductName string         `db:"product_name"`
	IsActive    bool           `db:"is_active"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}

type ProductRequest struct {
	ProductName string `json:"product_name" validate:"required,min=3,max=30"`
	IsActive    bool   `json:"is_active"`
}

type ProductVM struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type ProductRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []Product, total int, err error)
	Create(product Product) error
	Read(id string) (data Product, err error)
	Update(product Product) error
	Delete(product Product) error
	IsExist(id, name string) (isExist bool, err error)
}

type ProductUsecase interface {
	Browse(search, order, sort string, limit, page int) (productVM []ProductVM, pagination PaginationVM, err error)
	Create(req ProductRequest) (productVM ProductVM, err error)
	Read(id string) (productVM ProductVM, err error)
	Update(req ProductRequest, id string) (productVM ProductVM, err error)
	Delete(id string) (productVM ProductVM, err error)
	ConvertVM(product Product) ProductVM
}
