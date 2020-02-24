package main

import (
	"event_backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(e)

	//port := os.Getenv("PORT")

	// CORS policy
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(routes.Handlers())))
}
