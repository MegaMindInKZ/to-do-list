package main

import (
	"net/http"
)

func main() {
	config := NewConfig()

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: config.Address,
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
