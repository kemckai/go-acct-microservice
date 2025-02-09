package main

import (
	"fmt"
	"log"
	"net/http"

	"go-microservices/internal/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go Microservice!")
	})
	http.HandleFunc("/create", handlers.CreateAccount)
	http.HandleFunc("/get", handlers.GetAccount)
	http.HandleFunc("/update", handlers.UpdateAccount)
	http.HandleFunc("/delete", handlers.DeleteAccount)

	log.Println("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}