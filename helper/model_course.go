package helper

import (
	domain "courze-backend-app/model/domain"
	web "courze-backend-app/model/web"
)

func ToCourseResponse(course domain.Course) web.CourseResponse {
	return web.CourseResponse{
		Id:             course.Id,
		Name:           course.Name,
		Description:    course.Description,
		Price:          course.Price,
		Image:          course.Image,
		Video:          course.Video,
		InstructorName: course.InstructorName,
		SneakPeak:      course.SneakPeak,
	}
}

func ToCourseResponses(courses []domain.Course) []web.CourseResponse {
	var courseResponse []web.CourseResponse
	for _, course := range courses {
		courseResponse = append(courseResponse, ToCourseResponse(course))
	}
	return courseResponse
}
