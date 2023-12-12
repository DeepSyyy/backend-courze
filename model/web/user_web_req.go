package web

import "github.com/google/uuid"

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCourseRequest struct {
	UserId   uuid.UUID `json:"user_id"`
	CourseId int       `json:"course_id"`
}

type WishlistRequest struct {
	UserId   uuid.UUID `json:"user_id"`
	CourseId int       `json:"course_id"`
}
