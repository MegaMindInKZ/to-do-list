package config

import "flag"

type Config struct {
	Address    string
	Static     string
	dbusername string
	dbpassword string
	dbname     string
}

func NewConfig() Config {
	address := flag.String("address", "127.0.0.1:8080", "ip address and port number of our website")
	static := flag.String("static", "public", "static file storage")
	dbusername := flag.String("dbusername", "postgres", "database username")
	dbname := flag.String("dbname", "TO_DO_LIST", "database name")
	dbpassword := flag.String("dbpassword", "200103287sdu", "database password")
	flag.Parse()
	return Config{
		Address:    *address,
		Static:     *static,
		dbusername: *dbusername,
		dbpassword: *dbpassword,
		dbname:     *dbname,
	}
}
