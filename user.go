package main

type User struct {
	Id       int
	Username string
	Password string
}

func InsertUser(user *User) error {
	_, err := server.DB.Exec("INSERT INTO user(username, password_hash) VALUES(?, ?)", user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(id int) (*User, error) {
	row := server.DB.QueryRow("SELECT * FROM user WHERE id = ?", id)
	user := &User{}
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}
