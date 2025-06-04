package handlers

import (
	"net/http"
	"post_management/database"
	"post_management/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login handles user login and generates a JWT token
// It checks the provided email and password, and if valid, returns a JWT token.

func Login(c *gin.Context) {
	// Bind the incoming JSON to the AuthLogin struct
	var authLogin models.AuthLogin

	if err := c.ShouldBind(&authLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if the email Exists
	var userFound models.User
	database.DB.Where("email = ?", authLogin.Email).Find(&userFound)

	// If user not found, return an error
	// If userFound.ID is 0, it means the user does not exist
	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User Not Exist"})
		return
	}
	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authLogin.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong Password"})
		return
	}

	// If the email and password are valid, generate a JWT token
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with a secret key
	token, err := generateToken.SignedString([]byte("auth-api-jwt-secret"))

	// If there is an error signing the token, return an error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If the token is successfully generated, return it in the response
	c.JSON(http.StatusOK, gin.H{"token": token})

}
