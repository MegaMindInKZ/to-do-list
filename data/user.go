package data

import (
	"fmt"
	"strings"
	"time"
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

func UserByEmailOrUsername(usernameOrEmail string) (user User, err error) {
	if strings.Contains(usernameOrEmail, "@") {
		user, err = UserByEmail(usernameOrEmail)
	} else {
		user, err = UserByUsername(usernameOrEmail)
	}
	return
}

func UserByUsername(username string) (user User, err error) {
	err = DB.QueryRow("SELECT * FROM USERS WHERE USERNAME=$1", username).Scan(
		&user.ID, &user.UUID, &user.Username, &user.Name, &user.Email, &user.Password, &user.CreatedAt,
	)
	return
}

func UserByID(user_id int) (user User, err error) {
	err = DB.QueryRow("SELECT * FROM USERS WHERE ID = $1", user_id).Scan(
		&user.ID, &user.UUID, &user.Username, &user.Name, &user.Email, &user.Password, &user.CreatedAt,
	)
	return
}

func UserByEmail(email string) (user User, err error) {
	err = DB.QueryRow("SELECT * FROM USERS WHERE EMAIL = $1", email).Scan(
		&user.ID, &user.UUID, &user.Username, &user.Name, &user.Email, &user.Password, &user.CreatedAt,
	)
	fmt.Println(err)
	return
}
func (user *User) Create() (err error) {
	var existsUsername bool
	var existsEmail bool
	err = DB.QueryRow("SELECT EXISTS (SELECT EMAIL FROM USERS WHERE EMAIL=$1)", user.Email).Scan(&existsEmail)
	if err != nil {
		return
	}
	err = DB.QueryRow(
		"SELECT EXISTS (SELECT USERNAME FROM USERS WHERE USERNAME=$1)", user.Username,
	).Scan(&existsUsername)
	if err != nil {
		return
	}
	if existsUsername || existsEmail {
		//danger method
		return
	}
	st, err := DB.Prepare("INSERT INTO USERS(UUID, NAME, USERNAME, EMAIL, PASSWORD, CREATED_AT) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID, UUID, CREATED_AT")
	if err != nil {
		return
	}
	defer st.Close()
	err = st.QueryRow(CreateUUID(), user.Name, user.Username, user.Email, Encrypt(user.Password), time.Now()).Scan(
		&user.ID, &user.UUID, &user.CreatedAt,
	)
	return
}

func (user *User) Delete() (err error) {
	st, err := DB.Prepare("DELETE FROM USERS WHERE ID = ?")
	if err != nil {
		return
	}
	defer st.Close()
	_, err = st.Exec(user.ID)
	return
}

func (user *User) Update() (err error) {
	var existsUsername bool
	var existsEmail bool
	err = DB.QueryRow(
		"SELECT EXISTS (SELECT EMAIL FROM USERS WHERE EMAIL=$1 AND ID != $2)", user.Email, user.ID,
	).Scan(&existsEmail)
	if err != nil {
		//danger method
		return
	}
	err = DB.QueryRow(
		"SELECT EXISTS (SELECT USERNAME FROM USERS WHERE USERNAME=$1 AND ID != $2)", user.Username, user.ID,
	).Scan(&existsUsername)
	if err != nil {
		//danger method
		return
	}
	if existsUsername || existsEmail {
		//danger method
		return
	}
	stmt, err := DB.Prepare("UPDATE USERS SET NAME = $1, EMAIL = $2, USERNAME = $3 WHERE ID = $4")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Username, user.ID)
	return
}

func (user *User) CreateSession() (session Session, err error) {
	stmt, err := DB.Prepare("INSERT INTO SESSIONS (UUID, EMAIL, USER_ID, CREATED_AT) VALUES ($1, $2, $3, $4) RETURNING ID, UUID, EMAIL, USER_ID, CREATED_AT")
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(CreateUUID(), user.Email, user.ID, time.Now()).Scan(
		&session.ID, &session.UUID, &session.Email, &session.User_ID, &session.Created_at,
	)
	return
}

func (user *User) Session() (session Session, err error) {
	err = DB.QueryRow(
		"SELECT ID, UUID, EMAIL, USER_ID, CREATED_AT FROM SESSIONS WHERE USER_ID = ?", user.ID,
	).Scan(&session.ID, &session.UUID, &session.Email, &session.User_ID, &session.Created_at)
	return
}

// func (s Storage) InsertUser(user User) (err error) {

// }
