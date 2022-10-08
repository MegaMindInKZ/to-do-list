package main

import (
	"net/http"
	"to-do-list/data"
)

func login(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := database.UserByEmailOrUsername(request.PostFormValue("username-or-email"))
	if err != nil {

	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		return
	} else {
		http.Redirect(writer, request, "/login", 302)
	}

}
