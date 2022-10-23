package main

import (
	"html/template"
	"net/http"
	"to-do-list/data"
)

func index(writer http.ResponseWriter, request *http.Request) {
	h := indexStruct{}
	t, _ := template.ParseFiles("templates/base.html", "templates/index.html")
	t.ExecuteTemplate(writer, "base", h)
}

type indexStruct struct {
	User            data.User
	IsAuthenticated bool
}
