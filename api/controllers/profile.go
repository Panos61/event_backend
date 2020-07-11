package controllers

import (
	"encoding/json"
	"event_backend/api/auth"
	"event_backend/api/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//UpdateProfileData => Creates New Row or Updates Old Row
func (server *Server) UpdateProfileData(c *gin.Context) {

	// profile BEFORE UPDATE ENTITY
	profileBU := models.Profiles{}

	// Check if User is JWT Authenticated
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "User Unauthorized",
		})
		return
	}

	// Check if User exists
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	// Read Data
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Unable to get request (Read Data)",
		})
		return
	}

	err = json.Unmarshal(body, &profileBU)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot Unmarshal Body",
		})
		return
	}

	// Check if profile exists.If not, create a new row.
	// If not, update the row with the new data.

	profileBU.UserID = uid

	origProfile := models.Profiles{}
	err = server.DB.Debug().Model(models.Profiles{}).Where("user_id", uid).Take(&origProfile).Error
	if err != nil {
		log.Println("Error. The profile doesn't exist")
		profileBU.UserID = uid

		profileCreated, err := profileBU.SaveProfile(server.DB)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": profileCreated,
			})
			return
		}

	}

	profileCreated, err := profileBU.UpdateProfile(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": profileCreated,
	})

}

// GetMyProfile => Fetch my profile on request
func (server *Server) GetMyProfile(c *gin.Context) {
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	profile := models.Profiles{}

	profileSelected, err := profile.FindProfileByID(server.DB, uint64(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(200, profileSelected)
}

//GetProfile => Get A Specific Profile
// func (server *Server) GetProfile(c *gin.Context) {
// 	profileID := c.Param("id")
// 	pID, err := strconv.ParseUint(profileID, 10, 64)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  http.StatusBadRequest,
// 			"message": "Invalid Request",
// 		})
// 		return
// 	}

// 	profile := models.Profiles{}

// 	profileReceived, err := profile.FindProfileByID(server.DB, pID)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  http.StatusNotFound,
// 			"message": "No Profile Found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": profileReceived,
// 	})
// }

// GetProfileByID => Get specific profile based on ID
func (server *Server) GetProfileByID(c *gin.Context) {

	profileID := c.Param("user_id")

	pid, err := strconv.ParseInt(profileID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Profile Not Found.",
		})
		return
	}
	profile := models.Profiles{}

	profileSelected, err := profile.FindProfileByID(server.DB, uint64(pid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Profile Not Found.",
		})
		return
	}

	c.JSON(200, profileSelected)

}
