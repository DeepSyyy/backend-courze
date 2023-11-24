package service

import (
	"context"
	"courze-backend-app/model/domain"
	"courze-backend-app/model/web"
	"courze-backend-app/repository"
	"database/sql"
	"errors"
	"fmt"

	"courze-backend-app/helper"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)
	fmt.Println("pass1")
	hashPassword, err := helper.GenerateHashPassword(request.Password)
	if err != nil {
		helper.PanicIfError(err)
	}
	fmt.Println(hashPassword)
	fmt.Println("pass2")
	user := domain.User{
		Name:       request.Name,
		Email:      request.Email,
		Password:   hashPassword,
		IsVerified: false,
		CreatedAt:  helper.GetCurrentTime(),
	}
	fmt.Println("pass3")
	user, err = service.UserRepository.Register(ctx, tx, user)
	if err != nil {
		return web.UserResponse{}, err
	}
	fmt.Println("pass4")
	return helper.ToUserResponse(user), nil
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	user = service.UserRepository.UpdateUser(ctx, tx, user)

	return helper.ToUserResponse(user)

}

func (service *UserServiceImpl) LoginUser(ctx context.Context, request web.UserRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Email:    request.Email,
		Password: request.Password,
	}

	userFromDB, err := service.UserRepository.LoginUser(ctx, tx, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	// Verify password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(request.Password))
	if err != nil {
		return web.UserResponse{}, errors.New("invalid email or password")
	}

	return helper.ToUserResponse(userFromDB), nil
}

func (service *UserServiceImpl) GetUserByID(ctx context.Context, id string) (web.UserResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		// Tangani kesalahan pembukaan transaksi
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetUserByID(ctx, tx, id)
	if err != nil {
		// Tangani kesalahan pengambilan pengguna
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(user), nil
}
