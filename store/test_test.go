package store

import (
	"database/sql"
	"testing"
)

func getTestDb(t *testing.T) *sql.DB {
	db, err := MakeEmptyDatabase("../test.db", "../migration.sql")
	if err != nil {
		t.Fatalf("open database failed, %v", err.Error())
	}
	return db
}

func getTestStore(t *testing.T) *Store {
	db := getTestDb(t)
	return &Store{
		db: db,
	}
}

func TestInsertAndGetUser(t *testing.T) {
	s := getTestStore(t)
	defer s.db.Close()
	user1 := &User{
		Username: "yyc",
		Password: "123456",
	}
	user2 := &User{
		Username: "yyc2",
		Password: "123456",
	}
	s.InsertUser(user1)
	s.InsertUser(user2)
	if user1.Id == user2.Id {
		t.Error("get same id")
	}
	if got, err := s.GetUserById(user1.Id); err != nil || *got != *user1 {
		t.Error("GetUserById failed")
	}
	if got, err := s.GetUserByUsername(user2.Username); err != nil || *got != *user2 {
		t.Error("GetUserByUsername failed")
	}

	s.db.Close()
	s, _ = NewStore("../test.db")
	user3 := &User{
		Username: "yyc3",
		Password: "123456",
	}
	s.InsertUser(user3)
	if user3.Id != 3 {
		t.Error("got a wrong user id")
	}
}
