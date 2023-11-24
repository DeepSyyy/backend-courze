package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:"user_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"create_at"`
}

type ErrUserAlreadyExists struct {
	message string
}

func (e ErrUserAlreadyExists) Error() string {
	return "User already exists: " + e.message
}

func NewErrUserAlreadyExists(message string) *ErrUserAlreadyExists {
	return &ErrUserAlreadyExists{message}
}
