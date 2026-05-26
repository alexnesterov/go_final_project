package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexnesterov/go_final_project/internal/model"
	"github.com/alexnesterov/go_final_project/internal/service"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req model.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	id, err := service.CreateTask(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	fmt.Println("Created task with ID:", id)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"id": id})
}
