package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexnesterov/go_final_project/internal/domain/entity"
	"github.com/alexnesterov/go_final_project/internal/service"
)

func ReadTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")

	task, err := service.ReadTask(id)
	if errors.Is(err, entity.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(task)
}
