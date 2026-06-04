// Package handler
package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexnesterov/go_final_project/internal/service"
)

func NextDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	nowDate, err := parseNow(now)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	nextDate, err := service.NextDate(nowDate, date, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(nextDate))
}

func parseNow(n string) (time.Time, error) {
	if n == "" {
		return time.Now(), nil
	}

	now, err := time.Parse("20060102", n)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse now: %v", err)
	}

	return now, nil
}
