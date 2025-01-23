package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Auth AuthEmailConfig
}

type AuthEmailConfig struct {
	EmailUser    string
	EmailPass    string
	EmailAddress string
	EmailHost    string
	EmailPort    string
}

func NewConfig() *Config {
	err := godotenv.Load("3-validation-api/.env")
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Auth: AuthEmailConfig{
			EmailUser:    os.Getenv("EMAIL"),
			EmailPass:    os.Getenv("PASS"),
			EmailAddress: os.Getenv("ADDRESS"),
			EmailHost:    os.Getenv("EMAIL_HOST"),
			EmailPort:    os.Getenv("EMAIL_PORT"),
		},
	}
}
