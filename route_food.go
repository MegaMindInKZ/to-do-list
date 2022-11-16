package main

import (
	"html/template"
	"net/http"
	"to-do-list/data"
)

func foodMainPage(writer http.ResponseWriter, request *http.Request) {
	receipts, _ := data.AllReceipts()
	_, err := session(writer, request)
	if err == nil {
		t, _ := template.ParseFiles(
			"templates/base.html", "templates/private.navbar.html", "templates/food-main-page.html",
		)
		t.ExecuteTemplate(writer, "base", receipts)
	} else {
		t, _ := template.ParseFiles(
			"templates/base.html", "templates/public.navbar.html", "templates/food-main-page.html",
		)
		t.ExecuteTemplate(writer, "base", receipts)
	}
}

func receiptAddPage(writer http.ResponseWriter, request *http.Request) {

}
