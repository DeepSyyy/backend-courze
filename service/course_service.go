package service

import (
	"context"
	web "courze-backend-app/model/web"
)

type CourseService interface {
	GetAllCourse(ctx context.Context) []web.CourseResponse
	GetCourseById(ctx context.Context, courseId int) (web.CourseResponse, error)
	GetCourseByInstructorName(ctx context.Context, InstructorName string) []web.CourseResponse
	GetCourseByName(ctx context.Context, courseName string) ([]web.CourseResponse, error)
}
