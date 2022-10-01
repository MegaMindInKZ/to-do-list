package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello!!!")
}

func index(writer http.ResponseWriter, request *http.Request) {
	h := Hello{
		hello: "Hello World",
	}
	t, _ := template.ParseFiles("templates/base.html", "templates/index.html")
	t.ExecuteTemplate(writer, "base", h)
}

type Hello struct {
	hello string
}
