package service

import (
	"github.com/alexnesterov/go_final_project/internal/db"
	"github.com/alexnesterov/go_final_project/internal/domain/entity"
)

func ReadTask(id string) (*entity.Task, error) {
	return db.ReadTask(id)
}
