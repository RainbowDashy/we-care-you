package store

import (
	"context"
	"database/sql"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db  *sql.DB
	rdb *redis.Client
	ctx context.Context
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

func OpenRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

func NewStore(DBPath string) (*Store, error) {
	db, err := OpenDatabase(DBPath)
	if err != nil {
		return nil, err
	}
	rdb := OpenRedis()
	ctx := context.Background()

	return &Store{
		db:  db,
		rdb: rdb,
		ctx: ctx,
	}, nil
}
