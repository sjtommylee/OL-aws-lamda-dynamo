package common

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Init() (*sql.DB, error) {
	connStr := "user=postgres dbname=go-db sslmode=verify-full"
	driver := "postgres"
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
