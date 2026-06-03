package service

import (
	"fmt"

	"github.com/alexnesterov/go_final_project/internal/db"
	"github.com/alexnesterov/go_final_project/internal/model"
)

func CreateTask(req model.CreateTaskRequest) (string, error) {
	if req.Title == "" {
		return "", fmt.Errorf("%w: title is required", model.ErrCreateTask)
	}

	task := &model.Task{
		Date:    req.Date,
		Title:   req.Title,
		Comment: req.Comment,
		Repeat:  req.Repeat,
	}

	if err := checkDate(task); err != nil {
		return "", fmt.Errorf("%w: %w", model.ErrCreateTask, err)
	}

	id, err := db.CreateTask(task)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}
