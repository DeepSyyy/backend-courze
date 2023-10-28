package web

type CourseCreateRequest struct {
	Name         string   `json:"name" validate:"required"`
	Description  string   `json:"description" validate:"required"`
	Price        float64  `json:"price" validate:"required"`
	Image        string   `json:"image" validate:"required"`
	Video        string   `json:"video" validate:"required"`
	InstructorId int      `json:"instructor_id" validate:"required"`
	SneakPeak    []string `json:"sneak_peak" validate:"required"`
}
