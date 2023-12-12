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

func ToUserCourseResponse(usercourse domain.UserCourse) web.UserCourseResponse {
	return web.UserCourseResponse{
		Id:       usercourse.Id,
		UserId:   usercourse.UserId,
		CourseId: usercourse.CourseId,
	}
}

func ToUserCourseResponses(usercourses []domain.UserCourse) []web.UserCourseResponse {
	var userCourseResponse []web.UserCourseResponse
	for _, usercourse := range usercourses {
		userCourseResponse = append(userCourseResponse, ToUserCourseResponse(usercourse))
	}
	return userCourseResponse
}

func ToWishlistResponse(wishlist domain.Wishlist) web.WishlistResponse {
	return web.WishlistResponse{
		Id:       wishlist.Id,
		UserId:   wishlist.UserId,
		CourseId: wishlist.CourseId,
	}
}

func ToWishlistResponses(wishlists []domain.Wishlist) []web.WishlistResponse {
	var wishlistResponse []web.WishlistResponse
	for _, wishlist := range wishlists {
		wishlistResponse = append(wishlistResponse, ToWishlistResponse(wishlist))
	}
	return wishlistResponse
}
