package web

type CartCreateRequest struct {
	CourseId int `json:"course_id" validate:"required"`
	UserId   int `json:"user_id" validate:"required"`
}
