package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	CourseId  int       `json:"course_id"`
	Price     string    `json:"price"`
	Method    string    `json:"method"`
	CreatedAt time.Time `json:"create_at"`
}
