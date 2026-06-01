package service

import (
	"github.com/alexnesterov/go_final_project/internal/db"
	"github.com/alexnesterov/go_final_project/internal/domain/entity"
)

func ListTasks() ([]*entity.Task, error) {
	return db.ListTasks(50)
}
