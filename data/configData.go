package data

import (
	"fmt"
	"io/ioutil"
	"os"
)

func InitDB(database Storage) {
	st, ioErr := ioutil.ReadFile("data/setup.sql")
	if ioErr != nil {
		fmt.Println("Cannont read data/setup.sql")
		os.Exit(1)
	}
	database.Database.Exec(string(st))
}
