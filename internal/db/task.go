package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/alexnesterov/go_final_project/internal/domain/entity"
)

func CreateTask(task *entity.Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES ($1, $2, $3, $4)`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}

func ListTasks(limit int) ([]*entity.Task, error) {
	tasks := []*entity.Task{}

	query := `SELECT * FROM scheduler`
	args := []any{}

	if limit > 0 {
		query += ` LIMIT $1`
		args = append(args, limit)
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		task := &entity.Task{}

		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func ReadTask(id string) (*entity.Task, error) {
	task := &entity.Task{}

	query := `SELECT * FROM scheduler WHERE id = $1`

	row := DB.QueryRow(query, id)

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return task, nil
}

func UpdateTask(task *entity.Task) error {
	query := `
		UPDATE scheduler
		SET date = $1, title = $2, comment = $3, repeat = $4
		WHERE id = $5
	`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for updating task")
	}

	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = $1`

	res, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for deleting task")
	}

	return nil
}
