package db

import (
	"fmt"

	"github.com/alexnesterov/go_final_project/internal/model"
)

func CreateTask(task *model.Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES ($1, $2, $3, $4)`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}

func ListTasks(limit int) ([]*model.Task, error) {
	tasks := []*model.Task{}

	query := `SELECT * FROM scheduler`
	args := []any{}

	if limit > 0 {
		query += ` LIMIT $1`
		args = append(args, limit)
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("db query: %v", err)
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		task := &model.Task{}

		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, fmt.Errorf("db scan: %v", err)
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db rows: %v", err)
	}

	return tasks, nil
}
