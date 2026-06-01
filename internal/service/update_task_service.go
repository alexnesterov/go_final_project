package service

import (
	"fmt"

	"github.com/alexnesterov/go_final_project/internal/db"
	"github.com/alexnesterov/go_final_project/internal/domain/entity"
)

func UpdateTask(req entity.UpdateTaskRequest) error {
	if req.Title == "" {
		return fmt.Errorf("title is required")
	}

	task := &entity.Task{
		ID:      req.ID,
		Date:    req.Date,
		Title:   req.Title,
		Comment: req.Comment,
		Repeat:  req.Repeat,
	}

	if err := checkDate(task); err != nil {
		return err
	}

	return db.UpdateTask(task)
}
