package main

import (
	"log"
	"net/http"
	"time"
)

const PORT = "7540"

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", http.FileServer(http.Dir("web")).ServeHTTP)

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
