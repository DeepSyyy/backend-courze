package web

import "encoding/json"

type CourseCreateRequest struct {
	Name           string          `json:"name" validate:"required"`
	Description    string          `json:"description" validate:"required"`
	Price          float64         `json:"price" validate:"required"`
	Image          string          `json:"image" validate:"required"`
	Video          string          `json:"video" validate:"required"`
	InstructorName string          `json:"instructor_name" validate:"required"`
	SneakPeak      json.RawMessage `json:"sneak_peak" validate:"required"`
}
