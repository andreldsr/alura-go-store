package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connectionString := "user=hello dbname=hello password=hello host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("Connection error")
	}
	return db
}
