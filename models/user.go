package models

import (
	"time"
)

type User struct {
	ID             uint      `json:"id"`
	Username       string    `json:"username" validate:"required"`
	Email          string    `json:"email" validate:"required,email"`
	Password       string    `json:"password" validate:"required,min=6"`
	HashedPassword string    `json:"-"`
	Age            int       `json:"age" validate:"required,gte=8"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
