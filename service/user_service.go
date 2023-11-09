package service

import (
	"context"
	"courze-backend-app/model/web"
)

type UserService interface {
	Register(ctx context.Context, user web.UserRequest) web.UserResponse
	UpdateUser(ctx context.Context, user web.UserUpdateRequest) web.UserResponse
	LoginUser(ctx context.Context, user web.UserRequest) (web.UserResponse, error)
	GetUserByID(ctx context.Context, userID string) (web.UserResponse, error)
}
