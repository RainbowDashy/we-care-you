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
		if err := s.InsertItem(tx, item); err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}
