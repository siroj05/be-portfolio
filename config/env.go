package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWTSecret string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("env invalid")
	}
	JWTSecret = os.Getenv("JWT_SECRET")
}
