package service

import (
	"github.com/alexnesterov/go_final_project/internal/db"
	"github.com/alexnesterov/go_final_project/internal/model"
)

func ListTasks() ([]*model.Task, error) {
	return db.ListTasks(50)
}
