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
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-CSRF-Token", "Accept-Encoding", "Cache-Control",
		"X-header", "Access-Control-Allow-Methods", "x-access-token"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(routes.Handlers())))
}
