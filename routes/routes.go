package routes

import (
	"event_backend/controllers"
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

	//s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	s.HandleFunc("/user", controllers.GetLoggedUser).Methods("GET")
	return r
}

//CommonMiddleware => Set content-type
// func CommonMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
// 		next.ServeHTTP(w, r)
// 	})
// }
