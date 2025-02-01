package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbConf DbConfig
}

type DbConfig struct {
	DNS string
}

func GetConfig() Config {
	err := godotenv.Load("4-order-api/.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return Config{
		DbConf: DbConfig{
			DNS: os.Getenv("DB_DNS"),
		},
	}
}
