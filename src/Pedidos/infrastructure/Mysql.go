package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/pedidos_db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}
