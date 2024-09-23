package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
