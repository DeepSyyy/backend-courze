package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"user_id"`
	Name     string    `json:"nama"`
	Email    string    `json:"email"`
	Password float64   `json:"password"`
}
