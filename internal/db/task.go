package db

import "github.com/alexnesterov/go_final_project/internal/model"

func CreateTask(task *model.Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES ($1, $2, $3, $4)`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}
