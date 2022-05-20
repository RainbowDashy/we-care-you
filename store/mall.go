package store

import "database/sql"

type Mall struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"userid"`
	// BeginTime and EndTime are stored in format of unix timestamp
	BeginTime int64 `json:"begintime"`
	EndTime   int64 `json:"endtime"`
	// state == 0 means the mall is canceled
	// state == 1 means the mall is open
	State int64 `json:"state"`
}

type Item struct {
	Id          int64  `json:"id"`
	MallId      int64  `json:"mallid"`
	Total       int64  `json:"total"`
	Price       int64  `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

func (s *Store) InsertMall(tx *sql.Tx, mall *Mall) error {
	result, err := tx.Exec(`
		INSERT INTO mall(user_id, begin_time, end_time, state)
		VALUES(?, ?, ?, ?)
	`,
		mall.UserId,
		mall.BeginTime,
		mall.EndTime,
		mall.State,
	)
	if err != nil {
		return err
	}
	mall.Id, err = result.LastInsertId()
	return err
}

func (s *Store) UpdateMall(mall *Mall) error {
	_, err := s.db.Exec(`
		UPDATE mall
		SET begin_time = ?, end_time = ?, state = ?
		WHERE id = ?
	`, mall.BeginTime, mall.EndTime, mall.State, mall.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) InsertItem(tx *sql.Tx, item *Item) error {
	result, err := tx.Exec(`
		INSERT INTO item(mall_id, total, price, name, description, data)
		VALUES(?, ?, ?, ?, ?, ?)
	`,
		item.MallId, item.Total, item.Price, item.Name, item.Description, item.Data,
	)
	if err != nil {
		return err
	}
	item.Id, err = result.LastInsertId()
	return err
}

func (s *Store) CreateMall(user *User, mall *Mall, items []*Item) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
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

// scan order is Id, UserId, BeginTime, EndTime and State
func scanMallsFromRows(rows *sql.Rows) ([]*Mall, error) {
	malls := make([]*Mall, 0)
	for rows.Next() {
		mall := &Mall{}
		if err := rows.Scan(
			&mall.Id,
			&mall.UserId,
			&mall.BeginTime,
			&mall.EndTime,
			&mall.State,
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
		SELECT id, user_id, begin_time, end_time, state
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
		SELECT id, user_id, begin_time, end_time, state
		FROM mall
		WHERE user_id = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanMallsFromRows(rows)
}

func (s *Store) GetMallById(mallId int64) (*Mall, error) {
	rows, err := s.db.Query(`
		SELECT id, user_id, begin_time, end_time, state
		FROM mall
		WHERE id = ?
	`, mallId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	malls, err := scanMallsFromRows(rows)
	if err != nil {
		return nil, err
	}
	return malls[0], nil
}

func (s *Store) GetItemById(itemId int64) (*Item, error) {
	row := s.db.QueryRow(`
		SELECT id, mall_id, total, price, name, description, data
		FROM item
		WHERE id = ?
	`, itemId)

	item := &Item{}
	if err := row.Scan(
		&item.Id,
		&item.MallId,
		&item.Total,
		&item.Price,
		&item.Name,
		&item.Description,
		&item.Data,
	); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Store) GetItemsByMallId(mallId int64) ([]*Item, error) {
	rows, err := s.db.Query(`
		SELECT id, mall_id, total, price, name, description, data
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
			&item.Price,
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
