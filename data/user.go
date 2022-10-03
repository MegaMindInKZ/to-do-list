package data

type User struct {
	ID         int
	UUID       string
	Name       string
	Email      string
	Password   string
	Created_at string
}

func (s Storage) UserByID(user_id int) (user User, err error) {
	err = s.Database.QueryRow("SELECT * FROM USERS WHERE ID = ?", user_id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.Created_at)
	return
}

func (s Storage) UserByEmail(email string) (user User, err error) {
	err = s.Database.QueryRow("SELECT * FROM USERS WHERE EMAIL = ?", email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.Created_at)
	return
}

// func (s Storage) InsertUser(user User) (err error) {

// }
