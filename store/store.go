package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func OpenDatabase(DBPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBPath)
	return db, err
}

func NewStore(DBPath string) (*Store, error) {
	db, err := OpenDatabase(DBPath)
	if err != nil {
		return nil, err
	}
	return &Store{
		db: db,
	}, nil
}
