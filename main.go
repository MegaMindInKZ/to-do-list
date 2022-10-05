package main

import (
	"net/http"
	"to-do-list/config"
	"to-do-list/data"
)

func main() {
	config := config.NewConfig()
	database := data.NewStorage(config)
	data.InitDB(database)
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/sign-in", sign_in)

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
