package main

import (
	"errors"
	"net/http"
	"to-do-list/data"
)

func session(writer http.ResponseWriter, request *http.Request) (session data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		session = data.Session{UUID: cookie.Value}
		if ok, _ := session.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func isTrue(s string) bool {
	return s == "on"
}
