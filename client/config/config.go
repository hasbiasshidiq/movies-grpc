package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	API_PORT string
	GRPC_URL string
)

// LoadEnv Load Environment Variable
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		log.Println("using default environment variable")
	}

	API_PORT = os.Getenv("API_PORT")
	GRPC_URL = os.Getenv("GRPC_URL")

}
