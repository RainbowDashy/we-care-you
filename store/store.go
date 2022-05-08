package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(DBPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBPath)
	return db, err
}
