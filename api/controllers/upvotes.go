package controllers

import (
	"event_backend/api/auth"
	"event_backend/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpvotePost => ...
func (server *Server) UpvotePost(c *gin.Context) {

	postID := c.Param("id")
	// Check if post id is valid
	pid, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Request #1.",
		})
		return
	}
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	// check if user exists
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	// check if post exists
	post := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	upvote := models.Upvote{}
	upvote.UserID = user.ID
	upvote.PostID = post.ID

	upvoteCreated, err := upvote.SaveUpvote(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": upvoteCreated,
	})

}

// GetUpvotes => Get all upvotes from the post
func (server *Server) GetUpvotes(c *gin.Context) {

	postID := c.Param("id")

	// Check if post id is valid
	pid, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bard Request",
		})
		return
	}

	// Check if post exists
	post := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Post Not Found.",
		})
		return
	}

	upvote := models.Upvote{}

	upvotes, err := upvote.GetUpvotesInfo(server.DB, pid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Upvotes Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": upvotes,
	})
}

func (server *Server) RemoveUpvote(c *gin.Context) {

	upvoteID := c.Param("id")
	// check for valid post id
	pid, err := strconv.ParseUint(upvoteID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Unable to get request",
		})
		return
	}
	// check user auth
	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "User Unauthorized",
		})
		return
	}
	// Check if the post exist
	upvote := models.Upvote{}
	err = server.DB.Debug().Model(models.Upvote{}).Where("id = ?", pid).Take(&upvote).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "Post Not Found.",
		})
		return
	}
	// check if user is the owner of the post
	if uid != upvote.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  "User Unauthorized",
		})
		return
	}

	// remove the upvote
	_, err = upvote.DeleteUpvote(server.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "Not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Upvote Removed",
	})
}
