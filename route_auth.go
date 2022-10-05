package main

import (
	"net/http"
)

func login(_ http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	_, _ = database.UserByEmailOrUsername(request.PostFormValue("username-or-email"))
}
