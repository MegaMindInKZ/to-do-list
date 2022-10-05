package main

import (
	"net/http"
	"to-do-list/config"
	"to-do-list/data"
)

var database data.Storage

func main() {
	settings := config.NewConfig()
	database = data.NewStorage(settings)
	config.InitDB(database)
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(settings.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/sign-in", sign_in)
	mux.HandleFunc("/sign-up", sign_up)
	mux.HandleFunc("/login", login)

	server := &http.Server{
		Addr:    settings.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
