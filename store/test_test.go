package store

import (
	"context"
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
		db:  db,
		rdb: OpenRedis(),
		ctx: context.Background(),
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
	mall := &Mall{
		UserId:    user.Id,
		BeginTime: 0,
		EndTime:   10,
		State:     1,
	}
	if err := s.CreateMall(user, mall, items); err != nil {
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

func TestBuy(t *testing.T) {
	s := getTestStore(t)
	defer s.db.Close()
	seller := &User{
		Username: "seller",
		Password: "123456",
	}
	customer := &User{
		Username: "customer",
		Password: "123456",
	}
	s.InsertUser(seller)
	s.InsertUser(customer)
	items := []*Item{
		{Name: "A", Total: 10},
		{Name: "B", Total: 20},
		{Name: "C", Total: 30},
	}
	mall := &Mall{
		UserId:    seller.Id,
		BeginTime: 0,
		EndTime:   10,
		State:     1,
	}
	s.CreateMall(seller, mall, items)
	if err := s.Buy(customer, []*MallCustomer{
		{MallId: 1, UserId: customer.Id, ItemId: 1, BuyCount: 20},
	}); err == nil {
		t.Fatal("buycount is larger than total")
	}
	if err := s.Buy(customer, []*MallCustomer{
		{MallId: 1, UserId: customer.Id, ItemId: 1, BuyCount: 5},
		{MallId: 1, UserId: customer.Id, ItemId: 2, BuyCount: 5},
	}); err != nil {
		t.Fatal(err)
	}

	newItems, _ := s.GetItemsByMallId(1)
	for _, item := range newItems {
		if item.Name == "A" && item.Total != 5 ||
			item.Name == "B" && item.Total != 15 ||
			item.Name == "C" && item.Total != 30 {
			t.Error("wrong total")
		}
	}

	if orders, err := s.GetOrdersByItemId(1); err != nil {
		t.Fatal(err)
	} else {
		if len(orders) != 1 {
			t.Error("# of orders is wrong")
		}
		if orders[0].BuyCount != 5 {
			t.Error("wrong buy count")
		}
	}

	if orders, err := s.GetOrdersByUserId(customer.Id); err != nil {
		t.Fatal(err)
	} else {
		if len(orders) != 2 {
			t.Error("# of orders is wrong")
		}
		if orders[0].MallId != 1 {
			t.Error("wrong mallid")
		}
	}

	if orders, err := s.GetOrdersByMallId(1); err != nil {
		t.Fatal(err)
	} else {
		if len(orders) != 2 {
			t.Error("# of orders is wrong")
		}
		if orders[0].UserId != customer.Id {
			t.Error("wrong userid")
		}
	}
}
