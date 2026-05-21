package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexnesterov/go_final_project/internal/db"
	"github.com/alexnesterov/go_final_project/internal/handler"
)

const PORT = "7540"

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("init db: %v", err)
	}
	defer func() { _ = db.DB.Close() }()

	router := http.NewServeMux()

	router.HandleFunc("/", http.FileServer(http.Dir("web")).ServeHTTP)
	router.HandleFunc("/api/nextdate", handler.NextDate)

	server := http.Server{
		Addr:         ":" + PORT,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Server is running on port %s", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
