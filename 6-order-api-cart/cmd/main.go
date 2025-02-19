package main

import (
	"fmt"
	"godz/6-order-api-cart/configs"
	"godz/6-order-api-cart/internal/auth"
	"godz/6-order-api-cart/internal/user"
	"godz/6-order-api-cart/pkg/db"
	"godz/6-order-api-cart/pkg/middleware"

	"net/http"
)

func main() {
	conf := configs.GetConfig()
	dataBase := db.NewDB(conf)
	router := http.NewServeMux()

	userRepository := user.NewUserRepository(dataBase)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		UserRepo: userRepository,
		Config:   conf,
	})

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
