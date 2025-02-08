package main

import (
	"github.com/joho/godotenv"
	"godz/4-order-api/internal/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DNS")), &gorm.Config{})
	if err != nil {
		panic(err) // Лучше заменить на логирование и возврат ошибки
	}
	err = db.AutoMigrate(&product.Product{})
	if err != nil {
		panic(err)
	}
}
