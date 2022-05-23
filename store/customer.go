package store

import (
	"database/sql"
	"errors"
)

type MallCustomer struct {
	MallId   int64 `json:"mallid"`
	UserId   int64 `json:"userid"`
	ItemId   int64 `json:"itemid"`
	BuyCount int64 `json:"buycount"`
}

type Customer struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Location string `json:"location"`
	ItemId   int64  `json:"itemid"`
	MallId   int64  `json:"mallid"`
	BuyCount int64  `json:"buycount"`
}

func (s *Store) Buy(user *User, orders []*MallCustomer) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmtQuery, err := tx.Prepare(`
		SELECT mall_id, total
		FROM item
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer stmtQuery.Close()
	updateExec, err := tx.Prepare(`
		UPDATE item
		SET total = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer updateExec.Close()
	insertExec, err := tx.Prepare(`
		INSERT INTO mall_customer(mall_id, user_id, mall_item_id, buy_count)
		VALUES(?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer insertExec.Close()

	for _, order := range orders {
		var mallId, total int64
		if user.Id != order.UserId {
			return errors.New("userId does not match")
		}
		row := stmtQuery.QueryRow(order.ItemId)
		if err := row.Scan(&mallId, &total); err != nil {
			return err
		}
		if mallId != order.MallId {
			return errors.New("mallId does not match")
		}
		if order.BuyCount > total {
			return errors.New("buycount is larger than total count")
		}

		if _, err := updateExec.Exec(
			total-order.BuyCount,
			order.ItemId,
		); err != nil {
			return err
		}
		if _, err := insertExec.Exec(
			order.MallId,
			order.UserId,
			order.ItemId,
			order.BuyCount,
		); err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}

// Scan order is
//  MallId
// 	UserId
// 	ItemId
// 	BuyCount
func scanOrdersFromRows(rows *sql.Rows) ([]*MallCustomer, error) {
	orders := make([]*MallCustomer, 0)
	for rows.Next() {
		order := &MallCustomer{}
		if err := rows.Scan(
			&order.MallId,
			&order.UserId,
			&order.ItemId,
			&order.BuyCount,
		); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil

}

func (s *Store) GetOrdersByMallId(mallId int64) ([]*MallCustomer, error) {
	rows, err := s.db.Query(`
		SELECT mall_id, user_id, mall_item_id, buy_count
		FROM mall_customer
		WHERE mall_id = ?
	`, mallId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrdersFromRows(rows)
}

func (s *Store) GetOrdersByUserId(userId int64) ([]*MallCustomer, error) {
	rows, err := s.db.Query(`
		SELECT mall_id, user_id, mall_item_id, buy_count
		FROM mall_customer
		WHERE user_id = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrdersFromRows(rows)
}

func (s *Store) GetOrdersByItemId(itemId int64) ([]*MallCustomer, error) {
	rows, err := s.db.Query(`
		SELECT mall_id, user_id, mall_item_id, buy_count
		FROM mall_customer
		WHERE mall_item_id = ?
	`, itemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrdersFromRows(rows)
}

func (s *Store) GetCustomersByMallId(mallId int64) ([]*Customer, error) {
	rows, err := s.db.Query(`
		SELECT user.id, user.username, user.location, 
					 mall_customer.mall_id,
					 mall_customer.mall_item_id,
					 mall_customer.buy_count
		FROM user, mall_customer
		WHERE mall_customer.mall_id = ? AND user.id = mall_customer.user_id
	`, mallId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := make([]*Customer, 0)
	for rows.Next() {
		customer := &Customer{}
		if err := rows.Scan(
			&customer.Id,
			&customer.Username,
			&customer.Location,
			&customer.MallId,
			&customer.ItemId,
			&customer.BuyCount,
		); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}
