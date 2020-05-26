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

// CreateEvent => ..
func (server *Server) CreateEvent(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Unable to get request",
		})
		return
	}

	events := models.Events{}

	err = json.Unmarshal(body, &events)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot Unmarshal Body.",
		})
		return
	}

	uid, err := auth.ExtractTokenID(c.Request)
	//fmt.Println(uid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err,
		})
		return
	}

	// check if the user exists
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	events.CreatorID = uid

	eventCreated, err := events.SaveEvent(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": eventCreated,
	})

}

//GetEvents => Finds all Events created
func (server *Server) GetEvents(c *gin.Context) {
	event := models.Events{}

	events, err := event.FindAllEvents(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Events Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": events,
	})
}

// GetEvent => Fetch single Event
func (server *Server) GetEvent(c *gin.Context) {
	eventID := c.Param("id")
	eID, err := strconv.ParseUint(eventID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Request",
		})
		return
	}
	event := models.Events{}

	eventReceived, err := event.FindEventByID(server.DB, eID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Event Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": eventReceived,
	})
}

// GetUserEvents => Get all the user's events
func (server *Server) GetUserEvents(c *gin.Context) {
	userID := c.Param("id")

	uid, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Request",
		})
		return
	}

	event := models.Events{}
	events, err := event.FindUserEvents(server.DB, uint32(uid))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Event Found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": events,
	})
}
