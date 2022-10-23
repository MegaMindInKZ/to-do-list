package config

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Address    string
	Static     string
	DBusername string
	DBpassword string
	DBname     string
}

func NewConfig() Config {
	flag.Parse()
	return Config{
		Address:    *flag.String("address", "127.0.0.1:8080", "ip address and port number of our website"),
		Static:     *flag.String("static", "public", "static file storage"),
		DBusername: *flag.String("dbusername", "postgres", "database username"),
		DBname:     *flag.String("dbname", "to_do_list", "database name"),
		DBpassword: *flag.String("dbpassword", "200103287sdu", "database password"),
	}
}

func InitDB(database *sql.DB) {
	st, ioErr := ioutil.ReadFile("data/setup.sql")
	if ioErr != nil {
		fmt.Println("Cannont read data/setup.sql")
		os.Exit(1)
	}
	database.Exec(string(st))
}
