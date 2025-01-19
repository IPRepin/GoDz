package main

import (
	"godz/3-validation-api/config"
	"godz/3-validation-api/internal/verify"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{Config: cfg})
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
