package main

import (
	"net/http"
	"to-do-list/data"
)

func profileDeleteTask(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	err = request.ParseForm()
	if err != nil {
		//danger method
	}
}

func profileAddTask(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	err = request.ParseForm()
	if err != nil {
		//danger method
	}
	task := data.Task{
		Title:       request.PostFormValue("title"),
		UserID:      session.User_ID,
		Deadline:    request.PostFormValue("deadline"),
		Description: request.PostFormValue("description"),
		IsImportant: isTrue(request.PostFormValue("isImportant")),
	}
	task.Create()
	http.Redirect(writer, request, "/profile-tasks", 302)
}
