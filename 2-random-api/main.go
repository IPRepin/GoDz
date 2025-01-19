package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	router := http.NewServeMux()
	handler := &RandomNum{min: 0, max: 7}
	RandomHandler(router, handler)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Сервер слушает на :8081")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
