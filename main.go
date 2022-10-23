package main

import (
	"net/http"
	"to-do-list/data"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(data.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/sign-in", sign_in)
	mux.HandleFunc("/sign-up", sign_up)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/sign-up-account", signUpAccount)

	server := &http.Server{
		Addr:    data.Config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
