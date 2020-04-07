package profile

import (
	"encoding/json"
	"event_backend/api/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Server struct
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// CreateProfile => ..
func (server *Server) CreateProfile(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Unable to get request",
		})
		return
	}

	profile := models.Profile{}
	err = json.Unmarshal(body, &profile)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Cannot unmarshal body",
		})
		return
	}

	profile.Prepare()

	// Insert into DB
	profileCreated, err := profile.SaveProfile(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Error creating profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": profileCreated,
	})

}
