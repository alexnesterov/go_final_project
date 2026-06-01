package service

import (
	"time"

	"github.com/alexnesterov/go_final_project/internal/db"
)

func DoneTask(id string) error {
	task, err := db.ReadTask(id)
	if err != nil {
		return err
	}

	if task.Repeat == "" {
		return db.DeleteTask(id)
	}

	now := time.Now()

	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		return err
	}

	task.Date = next

	return db.UpdateTask(task)
}
