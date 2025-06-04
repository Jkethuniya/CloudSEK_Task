package handlers

import (
	"net/http"
	"post_management/database"
	"post_management/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CommentInput is the structure for the comment input data
type CommentInput struct {
	Content string `json:"content" binding:"required"`
}

// CreateComment handles the creation of a new comment on a post
func CreateComment(c *gin.Context) {
	// Retrieve the current user from the context
	user, _ := c.Get("currentUser")
	userData := user.(models.User)

	// Bind the incoming JSON data to the CommentInput struct
	var comment CommentInput
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIDStr := c.Param("post_id")
	postIDInt, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	// Convert the post ID to uint
	PostID := uint(postIDInt)

	// Check if the post exists
	var post []models.Post
	if err := database.DB.Model(&models.Post{}).Where("id = ?", PostID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// If the post does not exist, return a 404 error
	if len(post) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	// Create a new comment
	var cmt models.Comment
	cmt.UserID = userData.ID
	cmt.PostID = PostID
	cmt.Content = comment.Content

	// Save the comment to the database
	if err := database.DB.Create(&cmt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Return the created comment
	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// GetComments retrieves all comments for a specific post
func GetComments(c *gin.Context) {
	// Retrieve the post ID from the URL parameters
	postIDStr := c.Param("post_id")
	postIDInt, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Convert the post ID to uint
	PostID := uint(postIDInt)
	var post []models.Post
	if err := database.DB.Model(&models.Post{}).Where("id = ?", PostID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Check if the post exists
	if len(post) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Retrieve all comments for the specified post
	var comments []models.Comment
	if err := database.DB.Where("post_id = ?", PostID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	// If no comments are found, return a 404 error
	if len(comments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No comments found for this post"})
		return
	}

	// Return the comments

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func DeleteComment(c *gin.Context) {

	// Retrieve the current user from the context
	user, _ := c.Get("currentUser")
	userData := user.(models.User)

	// Find the comment ID from the URL parameters
	commentIDStr := c.Param("comment_id")
	commentIDInt, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	// Convert the comment ID to uint
	commentID := uint(commentIDInt)

	// Check if the comment exists
	var comment []models.Comment
	if err := database.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If the comment does not exist, return a 404 error
	if len(comment) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Check if the current user is the owner of the comment
	if comment[0].UserID != userData.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this comment"})
		return
	}

	// Delete the comment from the database
	if err := database.DB.Delete(&comment[0]).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
