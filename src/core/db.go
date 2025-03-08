package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// DB es una estructura que encapsula la conexión a la base de datos.
type DB struct {
	*sql.DB
}

// NewDB crea una nueva conexión a la base de datos MySQL.
func NewDB() (*DB, error) {
	db, err := sql.Open("mysql", "root:Desz117s@tcp(127.0.0.1:3306)/eventosdb")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return &DB{db}, nil
}
