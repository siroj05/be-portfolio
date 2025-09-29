package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var BaseUrlImg string

func LoadImgUrl() {
	err := godotenv.Load()
	if err != nil {
		log.Println("env invalid")
	}
	BaseUrlImg = os.Getenv("BASE_IMAGE_URL")
}
