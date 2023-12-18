package web

import "github.com/google/uuid"

type PaymentRequest struct {
	UserId   uuid.UUID `json:"user_id"`
	CourseId int       `json:"course_id"`
	Price    string    `json:"price"`
	Method   string    `json:"method"`
}
