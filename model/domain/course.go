package domain

type Course struct {
	CourseId          int     `json:"course_id"`
	CourseName        string  `json:"course_name"`
	CourseDescription string  `json:"course_desc"`
	CoursePrice       float64 `json:"course_price"`
	CourseImage       string  `json:"course_image"`
	CourseVideo       string  `json:"course_video"`
	InstructorId      int     `json:"instructor_id"`
}
