package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DbConnect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}
	return db, nil
}
