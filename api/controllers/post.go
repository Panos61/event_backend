package controllers

import (
	"encoding/json"
	"event_backend/api/auth"
	"event_backend/api/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreatePost(c *gin.Context) {

	// Read Data
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Unprocessable Entity",
		})
		return
	}

	post := models.Post{}

	err = json.Unmarshal(body, &post)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot Unmarshal body",
		})
		return
	}

	// Check if the user exists
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "User Unauthorized",
		})
		return
	}

	post.AuthorID = uid // Assign UserID as the AuthorID

	postCreated, err := post.SavePost(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": postCreated,
	})

	post.Prepare()

	// Error Validation
	errorVal := post.Validate()
	if len(errorVal) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Unprocessable Entity",
		})
		return
	}

}

func (server *Server) GetPosts(c *gin.Context) {
	post := models.Post{}

	posts, err := post.FindAllPosts(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Post Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": posts,
	})
}
