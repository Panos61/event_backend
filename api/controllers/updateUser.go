package controllers

import (
	"encoding/json"
	"event_backend/api/auth"
	"event_backend/api/models"
	"event_backend/api/security"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UpdatePassword => Update User's Password
func (server *Server) UpdatePassword(c *gin.Context) {
	// JSON Request Body
	requestBody := map[string]string{}

	userID := c.Param("id")

	// Checks if ID is valid
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID is not valid",
		})
		return
	}

	// Get the user ID from the token
	tokenID, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Error authenticating user",
		})
		return
	}

	if tokenID != 0 && tokenID != uint32(uid) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized User",
		})
		return
	}

	// IF AUTHENTICATION GOES WELL THEN UNMARSHAL BODY
	body, err := ioutil.ReadAll(c.Request.Body)
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot Unmarshal body",
		})
		return
	}

	// Before user update object
	beforeUserUpdate := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&beforeUserUpdate).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot find user",
		})
		return
	}

	// After user update object
	afterUserUpdate := models.User{}

	// ** VALIDATION **
	// If both password inputs and confirm input are empty
	if requestBody["password"] == "" && requestBody["newPassword"] == "" && requestBody["confirmPassword"] == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "No data given",
		})
		return
	}

	// If either old password OR new password OR confirm input is empty
	if requestBody["password"] == "" || requestBody["newPassword"] == "" || requestBody["confirmPassword"] == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Old password OR New password OR confirm input is empty",
		})
		return
	}

	if requestBody["newPassword"] != requestBody["confirmPassword"] {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Confirm password doesn't match with new password",
		})
		return
	}

	if requestBody["password"] != "" && requestBody["newPassword"] != "" && requestBody["confirmPassword"] != "" {

		// Validate old password
		err = security.VerifyPassword(beforeUserUpdate.Password, requestBody["password"])
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  http.StatusUnprocessableEntity,
				"message": "Current Password is wrong",
			})
			return
		}
		// ** END OF VALIDATION **

		// Update password
		afterUserUpdate.Password = requestBody["newPassword"]
		afterUserUpdate.Email = beforeUserUpdate.Email
		afterUserUpdate.Username = beforeUserUpdate.Username
		afterUserUpdate.Gender = beforeUserUpdate.Gender
	}

	afterUserUpdate.Prepare()

	updatedUser, err := afterUserUpdate.UpdateUserPassword(server.DB, uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": updatedUser,
	})
}

// UpdateEmail **
func (server *Server) UpdateEmail(c *gin.Context) {
	// JSON Request Body
	requestBody := map[string]string{}

	userID := c.Param("id")

	// Checks if ID is valid
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID is not valid",
		})
		return
	}

	// Get the user ID from the token
	tokenID, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Error authenticating user",
		})
		return
	}

	if tokenID != 0 && tokenID != uint32(uid) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized User",
		})
		return
	}

	// IF AUTHENTICATION GOES WELL THEN UNMARSHAL BODY
	body, err := ioutil.ReadAll(c.Request.Body)
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot Unmarshal body",
		})
		return
	}

	// Before user update object
	beforeUserUpdate := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&beforeUserUpdate).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot find user",
		})
		return
	}

	// After user update object
	afterUserUpdate := models.User{}

	// ** VALIDATION **
	// If both password inputs and confirm input are empty
	if requestBody["newEmail"] == "" && requestBody["password"] == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "No data given",
		})
		return
	}

	// If either new email or password is empty
	if requestBody["newEmail"] == "" || requestBody["password"] == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": " New email OR password input is empty",
		})
		return
	}

	if requestBody["newEmail"] != "" && requestBody["password"] != "" {

		// Validate password
		err = security.VerifyPassword(beforeUserUpdate.Password, requestBody["password"])
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  http.StatusUnprocessableEntity,
				"message": "Password is wrong",
			})
			return
		}
		// ** END OF VALIDATION **

		// Update email
		afterUserUpdate.Email = requestBody["newEmail"]
		afterUserUpdate.Password = beforeUserUpdate.Password
		afterUserUpdate.Username = beforeUserUpdate.Username
		afterUserUpdate.Gender = beforeUserUpdate.Gender
	}

	afterUserUpdate.Prepare()

	updatedUser, err := afterUserUpdate.UpdateUserEmail(server.DB, uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": updatedUser,
	})

}
