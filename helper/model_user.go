package helper

import (
	domain "courze-backend-app/model/domain"
	web "courze-backend-app/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponse []web.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, ToUserResponse(user))
	}
	return userResponse
}

func ToUserCourseResponses(usercourses []domain.UserCourse) []web.UserCourseResponse {
	var userCourseResponse []web.UserCourseResponse
	for _, usercourse := range usercourses {
		userCourseResponse = append(userCourseResponse, ToUserCourseResponse(usercourse))
	}
	return userCourseResponse
}

func ToUserCourseResponse(usercourse domain.UserCourse) web.UserCourseResponse {
	return web.UserCourseResponse{
		Id:       usercourse.Id,
		UserId:   usercourse.UserId,
		CourseId: usercourse.CourseId,
	}
}
