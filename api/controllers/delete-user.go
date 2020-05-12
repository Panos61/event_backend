package controllers

import (
	"event_backend/api/auth"
	"event_backend/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteUser => Deletes user's account
func (server *Server) DeleteUser(c *gin.Context) {
	var tokenID uint32
	userID := c.Param("id")

	// check is user id is valid
	uid, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User id is not valid",
		})
		return
	}

	tokenID, err = auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	if tokenID != 0 && tokenID != uint32(uid) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	user := models.User{}
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Error Finding User",
		})
	}

	// Delete user's events
	events := models.Events{}

	_, err = events.DeleteUserEvents(server.DB, uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error Deleting User",
		})
		return
	}

	// If no errors
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User Deleted",
	})
}
