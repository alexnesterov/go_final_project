package service

import (
	"fmt"
	"time"

	"github.com/alexnesterov/go_final_project/internal/db"
	"github.com/alexnesterov/go_final_project/internal/model"
)

func CreateTask(req model.CreateTaskRequest) (string, error) {
	if req.Title == "" {
		return "", fmt.Errorf("title is required")
	}

	task := model.Task{
		Date:    req.Date,
		Title:   req.Title,
		Comment: req.Comment,
		Repeat:  req.Repeat,
	}

	if err := checkDate(&task); err != nil {
		return "", err
	}

	id, err := db.CreateTask(&task)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(id), nil
}

func checkDate(task *model.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(`20060102`)
	}

	t, err := time.Parse(`20060102`, task.Date)
	if err != nil {
		return err
	}

	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
	}
	if err != nil {
		return err
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(`20060102`)
		} else {
			task.Date = next
		}
	}

	return nil
}
