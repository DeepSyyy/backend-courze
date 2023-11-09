package web

import (
	"encoding/json"
	"time"
)

type CourseResponse struct {
	Id             int             `json:"id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Price          float64         `json:"price"`
	Image          string          `json:"image"`
	Video          string          `json:"video"`
	InstructorName string          `json:"instructor_name"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	SneakPeak      json.RawMessage `json:"sneak_peak"`
}
