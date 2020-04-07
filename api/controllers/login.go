package controllers

import (
	"encoding/json"
	"event_backend/api/auth"
	"event_backend/api/models"
	"io/ioutil"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login => Unmarshals body / gets request
func (server *Server) Login(c *gin.Context) {

	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Unable to get request",
		})
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot unmarshal body",
		})
		return
	}
	errorMessages := user.Validate("login")
	if len(errorMessages) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": errorMessages,
		})
		return
	}

	user.Prepare()

	userData, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "User login error",
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"status": http.StatusOK,
	// 	// "user":   userData,
	// })
	c.Status(http.StatusOK)
	//json.NewEncoder(c.Writer).Encode(userData)
	c.JSON(200, userData)
}

// SignIn => Validates creds
func (server *Server) SignIn(email, password string) (map[string]interface{}, error) {
	var err error

	userData := make(map[string]interface{})

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		fmt.Println("Error getting user", user)
		return nil, err
	}

	// err = security.VerifyPassword(user.Password, password)
	// if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	// 	fmt.Println("Error hashing the password", err)
	// 	return nil, err
	// }

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil || errf == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("Error bcypt", errf)

	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		fmt.Println("Error creating token", err)
		return nil, err
	}

	userData["token"] = token
	// userData["id"] = user.ID
	// userData["email"] = user.Email
	// userData["username"] = user.Username
	// userData["gender"] = user.Gender
	// userData["password"] = user.Password
	userData["user"] = user

	return userData, nil

}
