package store

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	Id       int64
	Username string
	Password string
	Location string
}

// plaintext password -> hash
func NewUser(username, password, location string) *User {
	h := sha256.Sum256([]byte(password))
	return &User{
		Username: username,
		Password: fmt.Sprintf("%x", h),
		Location: location,
	}
}

func (s *Store) ValidUser(user *User) bool {
	userInDB, err := s.GetUserByUsername(user.Username)
	if err != nil {
		return false
	}
	return user.Password == userInDB.Password
}

func (s *Store) InsertUser(user *User) error {
	result, err := s.db.Exec(`
		INSERT INTO user(username, password_hash, location)
		VALUES(?, ?, ?)`,
		user.Username, user.Password, user.Location)
	if err != nil {
		return err
	}
	user.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserById(id int64) (*User, error) {
	row := s.db.QueryRow(`
		SELECT id, username, password_hash, location
		FROM user WHERE id = ?`, id)
	user := &User{}
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Location)
	return user, err
}

func (s *Store) GetUserByUsername(username string) (*User, error) {
	row := s.db.QueryRow(`
		SELECT id, username, password_hash, location
		FROM user WHERE username = ?`, username)
	user := &User{}
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Location)
	return user, err
}
