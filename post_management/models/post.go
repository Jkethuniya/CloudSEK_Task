package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`       // Unique identifier for the post
	UserID    uint      `json:"user_id"`                    // Foreign key to User
	Title     string    `json:"title" binding:"required"`   // Title of the post
	Content   string    `json:"content" binding:"required"` // Content of the post
	CreatedAt time.Time `json:"created_at"`                 // Timestamp when the post was created
	UpdatedAt time.Time `json:"updated_at"`                 // Timestamp when the post was last updated
}
