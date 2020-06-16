package api

import (
	"fmt"
	"log"
	"os"

	"event_backend/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}

// Run => Runs server
func Run() {
	var err error
	err = godotenv.Load()

	server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Check for .env load errors
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	fmt.Printf("Server up and running on port %s", apiPort)

	server.Run(apiPort)
}
