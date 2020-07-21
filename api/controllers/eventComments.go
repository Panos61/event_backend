package controllers

import (
	"encoding/json"
	"event_backend/api/auth"
	"event_backend/api/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UploadComment => Save and Upload Comment
func (server *Server) UploadComment(c *gin.Context) {
	eventID := c.Param("id")

	eid, err := strconv.ParseUint(eventID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Unable to handle request",
		})
		return
	}

	// Check the legitimacy of the token
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "User Unauthorized",
		})
		return
	}

	// Check if user exists
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "User not Found",
		})
		return
	}

	// Check if event exists
	event := models.Events{}
	err = server.DB.Debug().Model(models.Events{}).Where("id = ?", eid).Take(&event).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Event not Found",
		})
		return
	}

	eventComment := models.Event_Comment{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Unprocessable Entity",
		})
		return
	}

	err = json.Unmarshal(body, &eventComment)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot Unmarshal Body",
		})
		return
	}

	eventComment.UserID = uid
	eventComment.EventID = eid

	eventComment.Prepare()

	commentCreated, err := eventComment.SaveEventComment(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error while creating comment",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": commentCreated,
	})
}

// GetEventComments => fetches all event comments
func (server *Server) GetEventComments(c *gin.Context) {
	eventID := c.Param("id")

	eid, err := strconv.ParseUint(eventID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Unable to Handle Request",
		})
		return
	}

	// check if the event exists
	event := models.Events{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", eid).Take(&event).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Event Not Found",
		})
		return
	}

	comment := models.Event_Comment{}

	comments, err := comment.GetEventComments(server.DB, eid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Event Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": comments,
	})
}
