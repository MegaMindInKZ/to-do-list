package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"to-do-list/data"
)

func foodMainPage(writer http.ResponseWriter, request *http.Request) {
	receipts, err := data.AllReceipts()
	if err != nil {
		//danger method
		return
	}
	_, err = session(writer, request)
	if err == nil {
		t, _ := template.ParseFiles(
			"templates/base.html", "templates/private.navbar.html", "templates/food-main-page.html",
		)
		fmt.Println(t.ExecuteTemplate(writer, "base", receipts))
	} else {
		t, _ := template.ParseFiles(
			"templates/base.html", "templates/public.navbar.html", "templates/food-main-page.html",
		)
		t.ExecuteTemplate(writer, "base", receipts)
	}
}

func receiptAddPage(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	t, _ := template.ParseFiles("templates/base.html", "templates/private.navbar.html", "templates/food-add.html")
	t.ExecuteTemplate(writer, "base", nil)
}

func receiptAdd(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	request.ParseForm()
	var receipt data.Receipt
	receipt.User_ID = session.User_ID
	receipt.Name = request.PostFormValue("name")
	filename, err := pasteFileFood(request)
	if err == nil {
		receipt.Photo = filename
	}
	receipt.Duration, _ = strconv.Atoi(request.PostFormValue("duration"))
	receipt.Instruction = request.PostFormValue("instruction")
	err = receipt.Create()
	if err != nil {
		//danger method
		return
	}
	x, _ := strconv.Atoi(request.PostFormValue("ingredient-number"))
	for i := 0; i < x; i++ {
		var ingredient data.Ingredient
		ingredient.Name = request.PostFormValue("ingredient-name-" + strconv.Itoa(i))
		ingredient.ReceiptID = receipt.ID
		ingredient.Amount, _ = strconv.Atoi(request.PostFormValue("ingredient-amount-" + strconv.Itoa(i)))
		ingredient.Unit = request.PostFormValue("ingredient-unit-" + strconv.Itoa(i))
		err = ingredient.Create()
		if err != nil {
			fmt.Println(err)
			//danger method
			return
		}
	}
	http.Redirect(writer, request, "/receipts-main-page", 302)
}
