package main

import (
	"fmt"
	"godz/5-order-api-auth/configs"
	"godz/5-order-api-auth/internal/auth"
	"godz/5-order-api-auth/internal/user"
	"godz/5-order-api-auth/pkg/db"
	"godz/5-order-api-auth/pkg/middleware"
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
