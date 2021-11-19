package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/zeintkp/go-rest/domain"
)

func TestBrowse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	id1, _ := uuid.NewUUID()
	id2, _ := uuid.NewUUID()

	mockArticles := []domain.Product{
		{
			ID:          id1.String(),
			ProductName: "Product 1",
			IsActive:    true,
			CreatedAt:   time.Now().UTC().Format(time.RFC3339),
		},
		{
			ID:          id2.String(),
			ProductName: "Product 2",
			IsActive:    true,
			CreatedAt:   time.Now().UTC().Format(time.RFC3339),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "product_name", "is_active", "created_at", "updated_at", "deleted_at"}).
		AddRow(mockArticles[0].ID, mockArticles[0].ProductName, mockArticles[0].IsActive,
			mockArticles[0].CreatedAt, mockArticles[0].UpdatedAt.String, mockArticles[0].DeletedAt.String).
		AddRow(mockArticles[1].ID, mockArticles[1].ProductName, mockArticles[1].IsActive,
			mockArticles[1].CreatedAt, mockArticles[1].UpdatedAt.String, mockArticles[1].DeletedAt.String)

	search := "pro"
	order := "id"
	sort := "asc"
	limit := 10
	offset := 0

	query := `select * from "products" 
					where "product_name" ilike $1 
					and "deleted_at" is null 
					order by ` + order + ` ` + sort + ` limit $2 offset $3`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	rowTotal := sqlmock.NewRows([]string{"count"}).AddRow(2)

	query = `select count(*) from "products"
				where "product_name" ilike $1
				and "deleted_at" is null`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rowTotal)

	repo := NewProductRepository(db)
	data, total, err := repo.Browse(context.TODO(), search, order, sort, limit, offset)
	assert.Len(t, data, 2)
	assert.Equal(t, total, 2)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	now := time.Now()
	id := uuid.New().String()
	ar := domain.Product{
		ID:          id,
		ProductName: "test",
		IsActive:    true,
		CreatedAt:   now.UTC().Format(time.RFC3339),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `insert into "products" ("id","product_name", "is_active", "created_at") values ($1,$2,$3,$4)`

	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(ar.ID, ar.ProductName, ar.IsActive, ar.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewProductRepository(db)

	err = repo.Create(context.TODO(), ar)
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	id := uuid.New().String()

	rows := sqlmock.NewRows([]string{"id", "product_name", "is_active", "created_at", "updated_at", "deleted_at"}).
		AddRow(id, "title 1", true, time.Now(), time.Now(), nil)

	query := `select * from "products"
		where "id"=$1 and "deleted_at" is null`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	a := NewProductRepository(db)

	anArticle, err := a.Read(context.TODO(), id)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestUpdate(t *testing.T) {

	id := uuid.New().String()
	ar := domain.Product{
		ID:          id,
		ProductName: "test",
		IsActive:    true,
		UpdatedAt: sql.NullString{
			String: time.Now().UTC().Format(time.RFC3339),
			Valid:  true,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `update "products" set "product_name" = $1, "is_active" = $2, "updated_at" = $3 where "id" = $4`

	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(ar.ProductName, ar.IsActive, ar.UpdatedAt.String, ar.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewProductRepository(db)

	err = repo.Update(context.TODO(), ar)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {

	id := uuid.New().String()
	ar := domain.Product{
		ID: id,
		DeletedAt: sql.NullString{
			String: time.Now().UTC().Format(time.RFC3339),
			Valid:  true,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `update "products" set "deleted_at" = $1 where "id" = $2`

	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(ar.DeletedAt.String, ar.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewProductRepository(db)

	err = repo.Delete(context.TODO(), ar)
	assert.NoError(t, err)
}

func TestIsExist(t *testing.T) {

	id := uuid.New().String()
	productName := "test"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rowTotal := sqlmock.NewRows([]string{"count"}).AddRow(1)
	query := `select count(*) from "products" where lower("product_name") = $1 and "id" <> $2 and deleted_at is null`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rowTotal)

	rowTotal = sqlmock.NewRows([]string{"count"}).AddRow(0)
	query = `select count("id") from "products" where lower("product_name") = $1 and deleted_at is null`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rowTotal)

	repo := NewProductRepository(db)
	result, err := repo.IsExist(context.TODO(), id, productName)
	assert.Equal(t, true, result)
	assert.NoError(t, err)

	result, err = repo.IsExist(context.TODO(), "", productName)
	assert.Equal(t, false, result)
	assert.NoError(t, err)
}
