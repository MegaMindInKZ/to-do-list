package main

import "flag"

type Config struct {
	Address string
	Static  string
}

func NewConfig() Config {
	address := flag.String("address", "", "ip address and port number of our website")
	static := flag.String("static", "public", "static file storage")
	flag.Parse()
	return Config{
		Address: *address,
		Static:  *static,
	}
}
