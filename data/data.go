package data

import (
	"database/sql"
	"fmt"
	"os"
	"to-do-list/config"

	_ "github.com/lib/pq"
)

const (
	port = 5432
	host = "127.0.0.1"
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
		host, port, conf.DBusername, conf.DBpassword, conf.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("We have problems with connection database", err)
		os.Exit(1)
	}
	return Storage{
		Database: db,
	}
}
