package controllers

import (
	"event_backend/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FetchCityData => Fetch data from a specific city based on frontend url params
func (server *Server) FetchCityData(c *gin.Context) {
	eventParam := c.Param("city")

	event := models.Events{}
	events, err := event.CityEvents(server.DB, string(eventParam))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Events found for this location",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": events,
	})

}
