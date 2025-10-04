package main

import (
	httphandlers "backend/internal/infrastructure/http"
	"log"
	"net/http"
)

func main() {
	handler := httphandlers.NewRegistrationHandler()

	http.HandleFunc("/register", handler.Register)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
