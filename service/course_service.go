package service

import (
	"context"
	web "courze-backend-app/model/web"
)

type CourseService interface {
	GetAllCourse(ctx context.Context) []web.CourseResponse
	GetCourseById(ctx context.Context, courseId int) web.CourseResponse
	GetCourseByInstructorName(ctx context.Context, InstructorName string) []web.CourseResponse
}
