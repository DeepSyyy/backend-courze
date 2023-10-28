package domain

import "time"

type CourseCategory struct {
	CategoryId int       `json:"category_id"`
	CourseId   int       `json:"course_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Course     []Course  `json:"course"`
}
