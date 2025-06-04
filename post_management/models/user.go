package models

import (
	"time"
)

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey"`                            // Unique identifier for the user
	Name          string `json:"name"`                                            // Name of the user
	Password      string `json:"password"`                                        // Password of the user, should be hashed before storing
	Email         string `json:"email" gorm:"unique"`                             // Email of the user, must be unique
	ContactNumber string `json:"contact_number" binding:"required" gorm:"unique"` // Contact number of the user, must be unique
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type AuthInput struct {
	Name          string `json:"name" binding:"required"`                         // Name of the user, must not contain numbers
	Email         string `json:"email" binding:"required"`                        // Email of the user, must be unique
	Password      string `json:"password" binding:"required"`                     // Password of the user, must be at least 8 characters long and contain a special character
	ContactNumber string `json:"contact_number" binding:"required" gorm:"unique"` //
}

type AuthLogin struct {
	Email    string `json:"email" binding:"required"`    // Email of the user, must be unique
	Password string `json:"password" binding:"required"` // Password of the user, must be at least 8 characters long and contain a special character
}
