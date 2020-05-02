package controllers

import (
	"encoding/json"
	"event_backend/api/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// initProfile => ..
func (server *Server) initProfile(c *gin.Context) {
	profile := models.Profile{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Unable to get request",
		})
		return
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot unmarshal body",
		})
		return
	}

	//profile.Prepare()

	profileInitialized, err := profile.SaveProfile(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Error initializing profile",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      http.StatusCreated,
		"profileData": profileInitialized,
	})
}
