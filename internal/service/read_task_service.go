package service

import (
	"github.com/alexnesterov/go_final_project/internal/db"
	"github.com/alexnesterov/go_final_project/internal/model"
)

func ReadTask(id string) (*model.Task, error) {
	return db.ReadTask(id)
}
