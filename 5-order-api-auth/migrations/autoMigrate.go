package main

import (
	"github.com/joho/godotenv"
	"godz/5-order-api-auth/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DNS")), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		log.Println(err.Error())
	}
}
