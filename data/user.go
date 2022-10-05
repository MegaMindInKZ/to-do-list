package data

import (
	"strings"
)

type User struct {
	ID        int
	UUID      string
	Username  string
	Name      string
	Email     string
	Password  string
	CreatedAt string
}

func (s Storage) UserByEmailOrUsername(usernameOrEmail string) (user User, err error) {
	if strings.Contains(usernameOrEmail, "@") {
		user, err = s.UserByEmail(usernameOrEmail)
	} else {
		user, err = s.UserByUsername(usernameOrEmail)
	}
	return
}

func (s Storage) UserByUsername(username string) (user User, err error) {
	err = s.Database.QueryRow("SELECT * FROM USERS WHERE USERNAME = ?", username).Scan(&user.ID, &user.UUID, &user.Username, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (s Storage) UserByID(user_id int) (user User, err error) {
	err = s.Database.QueryRow("SELECT * FROM USERS WHERE ID = ?", user_id).Scan(&user.ID, &user.UUID, &user.Username, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (s Storage) UserByEmail(email string) (user User, err error) {
	err = s.Database.QueryRow("SELECT * FROM USERS WHERE EMAIL = ?", email).Scan(&user.ID, &user.UUID, &user.Username, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// func (s Storage) InsertUser(user User) (err error) {

// }
