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

func TestCreateAndGetMall(t *testing.T) {
	s := getTestStore(t)
	defer s.db.Close()
	user := &User{
		Username: "yyc",
		Password: "123456",
	}
	s.InsertUser(user)
	name := []string{"a", "b", "c"}
	total := []int64{1, 2, 3}
	description := []string{"A", "B", "C"}
	items := make([]*Item, 0, 3)
	for i := 0; i < 3; i += 1 {
		items = append(items, &Item{
			Name:        name[i],
			Total:       total[i],
			Description: description[i],
		})
	}
	if err := s.CreateMall(user, items); err != nil {
		t.Fatal(err)
	}
	malls, err := s.GetMallsByUserId(user.Id)
	if err != nil {
		t.Fatal(err)
	}
	if len(malls) != 1 {
		t.Error("wrong number of malls")
	}
	if malls[0].Id != 1 || malls[0].UserId != user.Id {
		t.Error("wrong mall")
	}
	gotItems, err := s.GetItemsByMallId(malls[0].Id)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range gotItems {
		flag := false
		for i := 0; i < 3; i += 1 {
			if item.Name == name[i] && item.Total == total[i] && item.Description == description[i] {
				flag = true
				break
			}
		}
		if !flag {
			t.Error("wrong items")
		}
	}
}
