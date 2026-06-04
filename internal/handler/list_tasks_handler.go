package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alexnesterov/go_final_project/internal/service"
)

func ListTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := service.ListTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("internal error: %v", err)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "internal error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]any{"tasks": tasks})
}
