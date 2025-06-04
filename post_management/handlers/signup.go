package handlers

import (
	"net/http"
	"post_management/database"
	"post_management/models"
	"post_management/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser handles user registration

func CreateUser(c *gin.Context) {
	var authInput models.AuthInput

	// Bind the incoming JSON to the AuthInput struct
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate Email
	// Check if the email is provided and not empty
	if !utils.IsValidEmail(authInput.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email format"})
		return
	}

	// Validate Contact Number
	if !utils.IsValidContactNumber(authInput.ContactNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid contact number"})
		return
	}

	// Validate Password
	if !utils.IsValidPassword(authInput.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters long and contain a special character"})
		return
	}
	var validatename string = authInput.Name
	var flag bool = true
	for i := 0; i < len(validatename); i++ {
		if validatename[i] >= '0' && validatename[i] <= '9' {
			flag = false
			break
		}
	}

	// Validate Name
	if flag == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name must not contain number"})
		return
	} else {

		// Check if the user already exists
		// If userFound.ID is not 0, it means the user already exists
		var userFound models.User
		database.DB.Where("email = ?", authInput.Email).Find(&userFound)

		if userFound.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User Already Exist"})
			return
		}

		// Hash the password
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create a new user in the database
		user := models.User{
			Name:          authInput.Name,
			Email:         authInput.Email,
			Password:      string(passwordHash),
			ContactNumber: authInput.ContactNumber,
		}
		authInput.Password = string(passwordHash)

		// Save the user to the database
		database.DB.Create(&user)
		c.JSON(http.StatusOK, gin.H{"data": authInput})
	}

}
