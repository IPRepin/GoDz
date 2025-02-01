package main

import (
	"fmt"
	"godz/4-order-api/configs"
	"godz/4-order-api/internal/product"
	"godz/4-order-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.GetConfig()
	database := db.NewDB(conf)
	router := http.NewServeMux()
	productRepo := product.NewProductRepo(database)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepo: productRepo,
	})
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("server is listening on :8081")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server error: %v\n", err)
	}
}
