package web

import "time"

type CourseUpdateRequest struct {
	Name           string    `json:"name" validate:"required"`
	Description    string    `json:"description" validate:"required"`
	Price          float64   `json:"price"`
	Image          string    `json:"image"`
	Video          string    `json:"video"`
	InstructorName string    `json:"instructor_id"`
	UpdatedAt      time.Time `json:"updated_at"`
	SneakPeak      []string  `json:"sneak_peak"`
}
