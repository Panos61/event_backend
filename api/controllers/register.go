package controllers

import (
	"encoding/json"
	//"event_backend/api/mail"

	"event_backend/api/auth"
	"event_backend/api/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	// Token, err := auth.CreateToken(user.ID)
	// if err != nil {
	// 	fmt.Println("Error creating token", err)
	// 	return
	// }
	// userData := make(map[string]interface{})
	// userData["token"] = Token

	// Send Welcome Email to the user
	//confirm, err := mail.SendMail.SendWelcomeMessage(user.Email, os.Getenv("SENDGRID_FROM"), user.Email, os.Getenv("SENDGRID_API_KEY"), os.Getenv("APP_ENV"))

	from := mail.NewEmail("Example User", "eventparkgr@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "panagiwtis.orovas@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": userCreated,
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

// GetMe => ..
func (server *Server) GetMe(c *gin.Context) {

	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
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
