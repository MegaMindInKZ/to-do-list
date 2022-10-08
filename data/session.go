package data

type Session struct {
	ID         int
	UUID       string
	Email      string
	User_ID    int
	Created_at string
}

func (session *Session) Check(db Storage) (valid bool, err error) {
	err = db.Database.QueryRow("SELECT ID, UUID, EMAIL, USER_ID, CREATED_AT FROM SESSIONS WHERE UUID = ?", session.UUID).Scan(&session.ID, &session.UUID, &session.Email, &session.User_ID, &session.Created_at)
	if err != nil {
		valid = false
		return
	}
	if session.ID != 0{
		valid = true
	}
	return
}

func (session *Session) DeleteByUUID(db Storage) (err error){
	statement := "delete from sessions where uuid = ?"
	stmt, err := db.Database.Prepare(statement)
	if err != nil{
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.UUID)
	return
}

func (session *Session) User(db Storage) (user User, err error){
	user = User{}
	err = db.
}