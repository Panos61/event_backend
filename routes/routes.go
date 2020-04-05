package routes

import (
	"event_backend/controllers"
	"event_backend/profile"
	"event_backend/utils/auth"

	"github.com/gorilla/mux"
)

// Handlers function
func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	//r.Use(CommonMiddleware)

	// Public routes
	r.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Auth routes
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)

	// Profile Routes
	//r.HandleFunc("/profile", profile.PostProfile).Methods("POST")
	r.HandleFunc("/profile", profile.FetchProfile).Methods("GET")
	r.HandleFunc("/profile", profile.UpdateProfile).Methods("PUT")

	s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	//s.HandleFunc("/user", controllers.GetLoggedInUser).Methods("GET")
	return r
}
