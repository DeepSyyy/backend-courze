package service

import (
	"context"
	"courze-backend-app/model/web"
)

type AdminService interface {
	CreateCourse(ctx context.Context, request web.CourseCreateRequest) web.CourseResponse
	UpdateCourse(ctx context.Context, request web.CourseUpdateRequest) web.CourseResponse
	DeleteCourse(ctx context.Context, courseID int)
}
