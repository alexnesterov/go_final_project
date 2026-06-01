// Package service
package service

import (
	"time"

	"github.com/alexnesterov/go_final_project/internal/domain/entity"
)

func checkDate(task *entity.Task) error {
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
