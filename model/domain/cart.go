package domain

type Cart struct {
	CartId   int `json:"cart_id"`
	CourseId int `json:"course_id"`
	UserId   int `json:"user_id"`
}
