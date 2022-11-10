package main

import (
	"errors"
	"io/ioutil"
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

func pasteFile(request *http.Request) (filename string, err error) {
	file, _, err := request.FormFile("avatar")
	if err != nil {
		return
	}
	defer file.Close()
	tempFile, err := ioutil.TempFile("private/avatar", "avatar-*.jpg")
	if err != nil {
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	tempFile.Write(fileBytes)
	filename = tempFile.Name()
	return
}
