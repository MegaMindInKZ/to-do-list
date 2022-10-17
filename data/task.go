package data

type Task struct {
	ID        int
	Title     string
	UserID    int
	CreatedAt string
}

func UserTasksByID(user_id int) (tasks []Task, err error) {
	rows, err := DB.Query("SELECT ID, TITLE, USER_ID, CREATED_AT FROM TASKS WHERE USER_ID = ?", user_id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Title, &task.UserID, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return
}

func DeleteUserTasks(user User) (err error) {
	stmt, err := DB.Prepare("DELETE FROM TASKS WHERE USER_ID = ?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	return
}

func (task *Task) Create() (err error) {
	st, err := DB.Prepare("INSERT INTO TASKS(TITLE, USER_ID, CREATED_AT) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	defer st.Close()
	err = st.QueryRow(task.Title, task.UserID, task.CreatedAt).Scan(&task.ID, &task.CreatedAt)
	return
}

func (task *Task) Delete() (err error) {
	stmt, err := DB.Prepare("DELETE FROM TASKS WHERE ID = ?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(task.ID)
	return
}

func (task *Task) Update(t Task) (err error) {
	stmt, err := DB.Prepare("UPDATE TASKS SET TITLE = ? WHERE ID = ?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(t.Title, task.ID)
	return
}
