package store

import (
	"crypto/sha256"
	"encoding/json"
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
	bytes, _ := json.Marshal(user)
	str := string(bytes)
	s.rdb.HSet(s.ctx, "user", fmt.Sprintf("%d", user.Id), str)
	s.rdb.HSet(s.ctx, "user", user.Username, str)
	return nil
}

func (s *Store) GetUserById(id int64) (*User, error) {
	result, err := s.rdb.HGet(s.ctx, "user", fmt.Sprintf("%d", id)).Result()
	if err == nil {
		user := &User{}
		err := json.Unmarshal([]byte(result), user)
		if err == nil {
			fmt.Println("Redis: cached")
			return user, nil
		}
	}

	row := s.db.QueryRow(`
		SELECT id, username, password_hash, location
		FROM user WHERE id = ?`, id)
	user := &User{}
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Location)
	return user, err
}

func (s *Store) GetUserByUsername(username string) (*User, error) {
	result, err := s.rdb.HGet(s.ctx, "user", username).Result()
	if err == nil {
		user := &User{}
		err := json.Unmarshal([]byte(result), user)
		if err == nil {
			fmt.Println("Redis: cached")
			return user, nil
		}
	}

	row := s.db.QueryRow(`
		SELECT id, username, password_hash, location
		FROM user WHERE username = ?`, username)
	user := &User{}
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Location)
	return user, err
}
