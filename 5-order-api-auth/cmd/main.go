package main

import (
	"fmt"
	"godz/5-order-api-auth/configs"
	"godz/5-order-api-auth/pkg/db"
	"godz/5-order-api-auth/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.GetConfig()
	_ = db.NewDB(conf)
	router := http.NewServeMux()

	stackMiddlewares := middleware.Chain(
		middleware.Cors,
		middleware.Logger,
	)

	server := &http.Server{
		Addr:    ":8081",
		Handler: stackMiddlewares(router),
	}
	fmt.Println("server is listening on :8081")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server error: %v\n", err)
	}
}
