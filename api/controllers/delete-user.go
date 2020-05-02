package controllers

import (
	"event_backend/api/auth"
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

}
