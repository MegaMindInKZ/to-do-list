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
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(writer, "Hello World")
}
