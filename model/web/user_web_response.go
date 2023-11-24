package web

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id         uuid.UUID `json:"user_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"create_at"`
}
