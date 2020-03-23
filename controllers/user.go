package controllers

import (
	"encoding/json"
	"event_backend/models"
	"event_backend/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// ErrorResponse ..
type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var db = utils.ConnectDB()

// Login => User Login
func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Stores user successful login into resp var
	resp := FindOne(user.Email, user.Password)

	json.NewEncoder(w).Encode(resp)
}

//FindOne => creds edit
func FindOne(email, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	// If no login error, generate a token and proceed
	// else return login error
	if errf == nil {

		tk := &models.Token{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: expiresAt,
			},
		}

		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}

		var resp = map[string]interface{}{"status": true, "message": "logged in"}
		resp["user"] = user
		resp["token"] = tokenString //Store the token in the response
		return resp
	} else {
		var resp = map[string]interface{}{"status": false, "message": "login error"}
		return resp
	}

}

//CreateUser function -- create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)

	createdUser := db.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//json.NewEncoder(w).Encode(createdUser)

	// Simple creds validation
	valErr := utils.ValidateUser(*user, utils.ValidationErrors)
	if len(valErr) > 0 {
		// db.AddError(errors.New("Wrong Validation Syntax"))
		fmt.Println("Wrong Val Syntax")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	// If no errors during registration, give create JWT
	if createdUser.Error == nil {

		//JWT implementation
		expiresAt := time.Now().Add(time.Minute * 100000).Unix()

		tk := &models.Token{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: expiresAt,
			},
		}

		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			log.Fatalln(err)
		}

		var resp = map[string]interface{}{"status": true, "message": "Signed up"}
		resp["token"] = tokenString //Store the token in the response
		resp["user"] = user

		json.NewEncoder(w).Encode(resp)

	} else {
		var resp = map[string]interface{}{"status": false, "message": "Registration failed"}
		json.NewEncoder(w).Encode(resp)
	}
}

//FetchUsers function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Preload("auth").Find(&users)

	json.NewEncoder(w).Encode(users)
}

// UpdateUser => Updates User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&user, id)
	json.NewDecoder(r.Body).Decode(user)
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

// DeleteUser => Deletes user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	db.Delete(&user)
	json.NewEncoder(w).Encode("User deleted")
}

// GetUser => Gets specific user
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	json.NewEncoder(w).Encode(&user)
}

// GetLoggedUser => Gets Logged In User with JWT
func GetLoggedUser(w http.ResponseWriter, r *http.Request) {

	//tk := models.Token{}

	var user models.User
	db.First(&user)

	json.NewEncoder(w).Encode(&user)

}
