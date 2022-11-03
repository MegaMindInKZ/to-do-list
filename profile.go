package main

import (
	"html/template"
	"net/http"
	"to-do-list/data"
)

func profileTasks(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	tasks, err := data.UserTasksByUserID(session.User_ID)
	t, err := template.ParseFiles(
		"templates/base.html", "templates/private.navbar.html",
		"templates/profile-base.html", "templates/profile-tasks.html",
	)
	t.ExecuteTemplate(writer, "base", tasks)
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
		Description: request.PostFormValue("description"),
		IsImportant: isTrue(request.PostFormValue("isImportant")),
	}
	task.Create()
	http.Redirect(writer, request, "/profile-tasks", 302)
}
func profileAddTaskPage(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	t, err := template.ParseFiles(
		"templates/base.html", "templates/private.navbar.html", "templates/profile-base.html",
		"templates/profile-task-add.html",
	)
	t.ExecuteTemplate(writer, "base", nil)
}
