package models

import (
	"time"
)

type SocialMedia struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name" validate:"required"`
	SocialMediaURL string    `json:"social_media_url" validate:"required"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
