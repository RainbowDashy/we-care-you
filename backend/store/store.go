package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data.db")
	return db, err
}
