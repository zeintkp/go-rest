package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/zeintkp/go-rest/domain"
)

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db: db}
}

type productRepository struct {
	db *sql.DB
}

func (repository *productRepository) Browse(ctx context.Context, search, order, sort string, limit, offset int) (data []domain.Product, total int, err error) {
	//hindari pake select *
	//validate args
	query := `select * from "products" 
					where "product_name" ilike $1 
					and "deleted_at" is null 
					order by ` + order + ` ` + sort + ` limit $2 offset $3`

	rows, err := repository.db.Query(query, "%"+search+"%", limit, offset)

	if err != nil {
		return data, total, err
	}

	dataTemp := domain.Product{}

	for rows.Next() { //row.Struct.scan
		rows.Scan(
			&dataTemp.ID,
			&dataTemp.ProductName,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)

		data = append(data, dataTemp)
	}

	if err = rows.Err(); err != nil {
		return data, total, err
	}

	query = `select count(*) from "products"
				where "product_name" ilike $1
				and "deleted_at" is null`
	err = repository.db.QueryRow(query, "%"+search+"%").Scan(&total)
	return data, total, err
}

func (repository *productRepository) Create(ctx context.Context, product domain.Product) error {
	query := `insert into "products" ("id","product_name", "is_active", "created_at") 
					values ($1,$2,$3,$4)`

	_, err := repository.db.Exec(query, product.ID, product.ProductName, product.IsActive, product.CreatedAt)
	return err
}

func (repository *productRepository) Read(ctx context.Context, id string) (data domain.Product, err error) {
	query := `select * from "products"
				where "id"=$1 and "deleted_at" is null`

	err = repository.db.QueryRow(query, id).Scan(
		&data.ID,
		&data.ProductName,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

func (repository *productRepository) Update(ctx context.Context, product domain.Product) error {
	query := `update "products" set 
				"product_name" = $1,
				"is_active" = $2,
				"updated_at" = $3
				where "id" = $4`

	_, err := repository.db.Exec(query, product.ProductName, product.IsActive, product.UpdatedAt.String, product.ID)
	return err
}

func (repository *productRepository) Delete(ctx context.Context, product domain.Product) error {
	query := `update "products" set 
				"deleted_at" = $1
				where "id" = $2`

	_, err := repository.db.Exec(query, product.DeletedAt.String, product.ID)
	return err
}

func (repository *productRepository) IsExist(ctx context.Context, id, name string) (isExist bool, err error) {
	var count int

	if id != "" {
		query := `select count(*) from "products" 
					where lower("product_name") = $1
					and "id" <> $2
					and deleted_at is null`
		err = repository.db.QueryRow(query, strings.ToLower(name), id).Scan(&count)
	} else {
		query := `select count("id") from "products" 
					where lower("product_name") = $1
					and deleted_at is null`
		err = repository.db.QueryRow(query, strings.ToLower(name)).Scan(&count)
	}

	return count > 0, err
}
