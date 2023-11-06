package web

type CourseResponse struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Price          float64  `json:"price"`
	Image          string   `json:"image"`
	Video          string   `json:"video"`
	InstructorName string   `json:"instructor_name"`
	SneakPeak      []string `json:"sneak_peak"`
}
