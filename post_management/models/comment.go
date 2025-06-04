package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PostID    uint      `json:"post_id" binding:"required"` // Foreign key to Post
	UserID    uint      `json:"user_id" binding:"required"` // Foreign key to User
	Content   string    `json:"content" binding:"required"` // Content of the comment
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
