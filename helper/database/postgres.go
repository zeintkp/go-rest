package database

import (
	"database/sql"
	"fmt"

	"github.com/zeintkp/go-rest/helper/config"
	"github.com/zeintkp/go-rest/helper/exception"

	_ "github.com/lib/pq" //pq is a pure Go Postgres driver for the database/sql package
)

//NewPostgresDatabase is used to create new Postgres setup
func NewPostgresDatabase(config config.Config) *sql.DB {

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?TimeZone=%s&sslmode=%s",
		config.Get("DB_USERNAME"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
		config.Get("DB_TIMEZONE"),
		config.Get("DB_SSL_MODE"),
	)

	if config.Get("DB_SSL_MODE") == "require" {
		connStr += fmt.Sprintf("&sslcert=%s&sslkey=%s&sslrootcert=%s",
			config.Get("DB_SSL_CERT"),
			config.Get("DB_SSL_KEY"),
			config.Get("DB_SSL_ROOT_CERT"),
		)
	}

	db, err := sql.Open("postgres", connStr)
	exception.PanicIfNeeded(err)

	return db
}
