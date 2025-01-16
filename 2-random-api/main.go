package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	RandomHandler(router)
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("server is listening on :8081")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server error: %v\n", err)
	}
}
