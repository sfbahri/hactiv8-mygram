package models

import (
	"time"
)

type Photo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title" validate:"required"` // Title must not be empty
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" validate:"required"` // Photo URL must not be empty
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
