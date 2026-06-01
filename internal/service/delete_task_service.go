package service

import "github.com/alexnesterov/go_final_project/internal/db"

func DeleteTask(id string) error {
	return db.DeleteTask(id)
}
