package store

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func MakeEmptyDatabase(DBPath string, migrationPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		return nil, err
	}
	flag := false
	defer func() {
		if flag {
			db.Close()
		}
	}()
	bytes, err := os.ReadFile(migrationPath)
	if err != nil {
		flag = true
		return nil, err
	}
	sqlStmts := string(bytes)
	_, err = db.Exec(sqlStmts)
	if err != nil {
		flag = true
		return nil, err
	}
	return db, nil
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
