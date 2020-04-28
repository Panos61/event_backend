package controllers

import (
	"encoding/json"
	"event_backend/api/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEvent => ..
func (server *Server) CreateEvent(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"message": "Unable to get request"
		})
		return
	}

	events := models.Events{}

	err = json.Unmarshal(body, &events)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"message": "Cannot Unmarshal Body."
		})
		return
	}
}