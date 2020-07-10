package controllers

import (
	"event_backend/api/auth"
	"event_backend/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// initProfile => ..
// func (server *Server) initProfile(c *gin.Context) {
// 	profile := models.Profiles{}

// 	body, err := ioutil.ReadAll(c.Request.Body)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"status":  http.StatusUnprocessableEntity,
// 			"message": "Unable to get request",
// 		})
// 		return
// 	}

// 	err = json.Unmarshal(body, &profile)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"status":  http.StatusUnprocessableEntity,
// 			"message": "Cannot unmarshal body",
// 		})
// 		return
// 	}

// 	uid, err := auth.ExtractTokenID(c.Request)
// 	//fmt.Println(uid)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"status":  http.StatusUnauthorized,
// 			"message": err,
// 		})
// 		return
// 	}

// 	// check if the user exists
// 	user := models.User{}
// 	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error

// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"status":  http.StatusUnauthorized,
// 			"message": "Unauthorized",
// 		})
// 		return
// 	}

// 	profile.UserID = uid

// 	profileCreated, err := profile.SaveProfile(server.DB)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  http.StatusInternalServerError,
// 			"message": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"status":  http.StatusCreated,
// 		"message": profileCreated,
// 	})
// }

func (server *Server) updateProfileData(c *gin.Context) {
	// JSON Request Body
	requestBody := map[string]string{}

	userID := c.Param("id")

	// checks if user ID is valid
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

	// Before profile update object
	//beforeProfileUpdate := models.Profiles{}

	// err = server.DB.Debug().Model(models.Profiles{}).Where("id = ?", uid).Take(&beforeProfileUpdate).Error
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"status":  http.StatusUnprocessableEntity,
	// 		"message": "Cannot find user",
	// 	})
	// 	return
	// }

	// After profile update object
	afterProfileUpdate := models.Profiles{}

	// Update Profile Data
	afterProfileUpdate.FirstName = requestBody["firstName"]
	afterProfileUpdate.LastName = requestBody["lastName"]
	afterProfileUpdate.Introduction = requestBody["introduction"]
	afterProfileUpdate.Age = requestBody["age"]

	updatedProfile, err := afterProfileUpdate.UpdateProfile(server.DB, uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": updatedProfile,
	})
}

//GetProfiles => Finds All Profiles
func (server *Server) GetProfiles(c *gin.Context) {
	profile := models.Profiles{}

	profiles, err := profile.FindAllProfiles(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Profile Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": profiles,
	})
}

//GetProfile => Get A Specific Profile
func (server *Server) GetProfile(c *gin.Context) {
	profileID := c.Param("id")
	pID, err := strconv.ParseUint(profileID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Request",
		})
		return
	}

	profile := models.Profiles{}

	profileReceived, err := profile.FindProfileByID(server.DB, pID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Profile Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": profileReceived,
	})
}
