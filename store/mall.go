package store

import "database/sql"

type Mall struct {
	Id     int64
	UserId int64
}

type Item struct {
	Id          int64
	MallId      int64
	Total       int64
	Name        string
	Description string
	Data        string
}

func (s *Store) InsertMall(tx *sql.Tx, mall *Mall) error {
	result, err := tx.Exec("INSERT INTO mall(user_id) VALUES(?)", mall.UserId)
	if err != nil {
		return err
	}
	mall.Id, err = result.LastInsertId()
	return err
}

func (s *Store) InsertItem(tx *sql.Tx, item *Item) error {
	result, err := tx.Exec(`
		INSERT INTO item(mall_id, total, name, description, data)
		VALUES(?, ?, ?, ?, ?)
	`,
		item.MallId, item.Total, item.Name, item.Description, item.Data,
	)
	if err != nil {
		return err
	}
	item.Id, err = result.LastInsertId()
	return err
}

func (s *Store) CreateMall(user *User, items []*Item) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	mall := &Mall{UserId: user.Id}
	if err := s.InsertMall(tx, mall); err != nil {
		return err
	}
	for _, item := range items {
		item.MallId = mall.Id
		if err := s.InsertItem(tx, item); err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}

func scanMallsFromRows(rows *sql.Rows) ([]*Mall, error) {
	malls := make([]*Mall, 0)
	for rows.Next() {
		mall := &Mall{}
		if err := rows.Scan(
			&mall.Id,
			&mall.UserId,
		); err != nil {
			return nil, err
		}
		malls = append(malls, mall)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return malls, nil
}

func (s *Store) GetMalls() ([]*Mall, error) {
	rows, err := s.db.Query(`
		SELECT id, user_id
		FROM mall
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanMallsFromRows(rows)
}

func (s *Store) GetMallsByUserId(userId int64) ([]*Mall, error) {
	rows, err := s.db.Query(`
		SELECT id, user_id
		FROM mall
		WHERE user_id = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanMallsFromRows(rows)
}

func (s *Store) GetItemsByMallId(mallId int64) ([]*Item, error) {
	rows, err := s.db.Query(`
		SELECT id, mall_id, total, name, description, data
		FROM item
		WHERE mall_id = ?
	`, mallId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*Item, 0)
	for rows.Next() {
		item := &Item{}
		if err := rows.Scan(
			&item.Id,
			&item.MallId,
			&item.Total,
			&item.Name,
			&item.Description,
			&item.Data,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
