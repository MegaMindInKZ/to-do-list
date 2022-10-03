package data

import (
	"database/sql"
	"fmt"
	"os"
	"to-do-list/config"

	_ "github.com/lib/pq"
)

const (
	port     = 5432
	host     = "127.0.0.1"
	user     = "zanggar"
	dbname   = "postgres"
	password = "200103287sdu"
)

type Session struct {
	ID         int
	UUID       string
	Email      string
	User_ID    int
	Created_at string
}

type Task struct {
	ID      int
	Title   string
	User_ID int
}

type Storage struct {
	Database *sql.DB
}

func NewStorage(conf config.Config) Storage {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("We have problems with connection database", err)
		os.Exit(1)
	}
	var s int
	err = db.QueryRow("select id from test").Scan(&s)
	fmt.Println(err)
	fmt.Println(s)
	return Storage{
		Database: db,
	}
}
