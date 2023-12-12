package service

import (
	"context"
	"courze-backend-app/model/web"
)

type UserService interface {
	Register(ctx context.Context, user web.UserRequest) (web.UserResponse, error)
	UpdateUser(ctx context.Context, user web.UserUpdateRequest) web.UserResponse
	LoginUser(ctx context.Context, user web.UserRequest) (web.UserResponse, error)
	GetUserByID(ctx context.Context, userID string) (web.UserResponse, error)
	Enroll(ctx context.Context, usercourse web.UserCourseRequest) (web.UserCourseResponse, error)
	GetUserCourseByID(ctx context.Context, userID string) []web.UserCourseResponse
	AddWishlist(ctx context.Context, usercourse web.WishlistRequest) (web.WishlistResponse, error)
	GetWishlistByID(ctx context.Context, userID string) []web.WishlistResponse
}
