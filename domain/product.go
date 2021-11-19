package domain

import (
	"context"
	"database/sql"
)

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
	Browse(ctx context.Context, search, order, sort string, limit, offset int) (data []Product, total int, err error)
	Create(ctx context.Context, product Product) error
	Read(ctx context.Context, id string) (data Product, err error)
	Update(ctx context.Context, product Product) error
	Delete(ctx context.Context, product Product) error
	IsExist(ctx context.Context, id, name string) (isExist bool, err error)
}

type ProductUsecase interface {
	Browse(ctx context.Context, search, order, sort string, limit, page int) (productVM []ProductVM, pagination PaginationVM, err error)
	Create(ctx context.Context, req ProductRequest) (productVM ProductVM, err error)
	Read(ctx context.Context, id string) (productVM ProductVM, err error)
	Update(ctx context.Context, req ProductRequest, id string) (productVM ProductVM, err error)
	Delete(ctx context.Context, id string) (productVM ProductVM, err error)
	ConvertVM(product Product) ProductVM
}
