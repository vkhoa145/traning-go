package application

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vkhoa145/go-training/app/server"
	"github.com/vkhoa145/go-training/config"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.LoadConfig()
	server := server.NewServer(config)
	error := server.Start()
	if error != nil {
		log.Fatal("Error starting server: ", error)
	}
}
