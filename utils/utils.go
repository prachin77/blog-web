package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort int
	ClientPort int
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load("P:/blog-web/.env"); err != nil {
		return nil, errors.New("error loading .env file")
	}

	// Get server port
	serverPortStr := os.Getenv("SERVER_PORT")
	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil || serverPort <= 0 {
		return nil, errors.New("invalid SERVER_PORT in .env file")
	}

	// Get client port
	clientPortStr := os.Getenv("CLIENT_PORT")
	clientPort, err := strconv.Atoi(clientPortStr)
	if err != nil || clientPort <= 0 {
		return nil, errors.New("invalid CLIENT_PORT in .env file")
	}

	return &AppConfig{
		ServerPort: serverPort,
		ClientPort: clientPort,
	}, nil
}