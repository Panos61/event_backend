package controllers

import (
	"encoding/json"
	"event_backend/api/models"
	"event_backend/api/security"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser => Creates new user
func (server *Server) CreateUser(c *gin.Context) {

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

	// Password Hash and Salt ** BCRYPT
	pass, err := security.Hash(user.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Error hashing password",
		})
		return
	}

	user.Password = string(pass)

	// *** ***

	user.Prepare()

	// Insert user into DB
	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error creating user",
		})
		return
	}

	// token, err := auth.CreateToken(user.ID)
	// if err != nil {
	// 	fmt.Println("Error creating token", err)
	// 	return
	// }
	// userData := make(map[string]interface{})
	// userData["token"] = token

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": userCreated,
		//"token":    userData,
	})
}

// GetUsers => Gets All Users
func (server *Server) GetUsers(c *gin.Context) {

	errList = map[string]string{}

	user := models.User{}

	users, err := user.FindUsers(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Users not found",
		})
	}

	c.JSON(200, users)
}

// GetUserByID => ..
func (server *Server) GetUserByID(c *gin.Context) {

	errList = make(map[string]string)

	userID := c.Param("id")

	uid, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User not found",
		})
		return
	}
	user := models.User{}

	userSelected, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(200, userSelected)
}
