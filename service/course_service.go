package service

import (
	"context"
	web "courze-backend-app/model/web"
)

type CourseService interface {
	CreateCourse(ctx context.Context, request web.CourseCreateRequest) web.CourseResponse
	GetAllCourse(ctx context.Context) []web.CourseResponse
}
