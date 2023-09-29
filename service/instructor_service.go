package service

import (
	"context"
	"courze-backend-app/model/web"
)

type InstructorService interface {
	Insert(ctx context.Context, request web.InstructorCreateRequest) web.InstructorResponse
}
