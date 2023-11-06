package domain

type Course struct {
	Id             int      `json:"course_id"`
	Name           string   `json:"course_name"`
	Description    string   `json:"course_desc"`
	Price          float64  `json:"course_price"`
	Image          string   `json:"course_image"`
	Video          string   `json:"course_video"`
	InstructorName string   `json:"instructor_name"`
	SneakPeak      []string `json:"sneak_peak"`
}
