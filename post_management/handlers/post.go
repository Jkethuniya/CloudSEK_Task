package handlers

import (
	"net/http"
	"post_management/database"
	"post_management/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePost handles the creation of a new post
// It expects the post data in the request body and associates it with the current user.
func CreatePost(c *gin.Context) {
	user, _ := c.Get("currentUser")
	userData := user.(models.User)
	var post models.Post

	// Bind the incoming JSON data to the Post struct
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set the UserID of the post to the current user's ID
	post.UserID = userData.ID
	database.DB.Create(&post)

	// Return the created post
	c.JSON(http.StatusOK, gin.H{"data": post})

}

// GetPosts retrieves all posts from the database
func GetPosts(c *gin.Context) {
	var posts []models.Post
	// Retrieve all posts from the database
	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	// If no posts are found, return a 404 Not Found response

	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No posts found"})
		return
	}

	// Return the list of posts
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// DeletePost handles the deletion of a post by its ID
func DeletePost(c *gin.Context) {
	user, _ := c.Get("currentUser")
	userData := user.(models.User)
	var post models.Post
	// Retrieve the post ID from the URL parameter
	postIDStr := c.Param("id")
	// fmt.Println("Post ID:", postID)
	postIDInt, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert the post ID to uint
	postID := uint(postIDInt)

	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if the post exists and belongs to the current user
	if post.UserID != userData.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this post"})
		return
	}

	// If the post exists, delete it and its associated comments
	if err := database.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	// Delete all comments associated with the post
	if err := database.DB.Where("post_id = ?", postID).Delete(&models.Comment{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associated comments"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
