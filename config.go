package main

import "flag"

type Config struct {
	Address string
	Static  string
}

func NewConfig() Config {
	return Config{
		Address: *flag.String("address", "", "host of our website"),
		Static:  *flag.String("static", "", ""),
	}
}
