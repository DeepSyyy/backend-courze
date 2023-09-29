package helper

import (
	domain "courze-backend-app/model/domain"
	web "courze-backend-app/model/web"
)

func ToCourseResponse(course domain.Course) web.CourseResponse {
	return web.CourseResponse{
		Id:           course.CourseId,
		Name:         course.CourseName,
		Description:  course.CourseDescription,
		Price:        course.CoursePrice,
		Image:        course.CourseImage,
		Video:        course.CourseVideo,
		InstructorId: course.InstructorId,
	}
}

func ToCourseResponses(courses []domain.Course) []web.CourseResponse {
	var courseResponse []web.CourseResponse
	for _, course := range courses {
		courseResponse = append(courseResponse, ToCourseResponse(course))
	}
	return courseResponse
}
