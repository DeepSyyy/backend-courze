package service

import (
	"context"
	"courze-backend-app/model/domain"
	"courze-backend-app/model/web"
	repository "courze-backend-app/repository"
	"database/sql"

	"courze-backend-app/helper"

	"github.com/go-playground/validator"
)

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewAdminService(adminRepository repository.AdminRepository, DB *sql.DB, validate *validator.Validate) AdminService {
	return &AdminServiceImpl{
		AdminRepository: adminRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *AdminServiceImpl) CreateCourse(ctx context.Context, request web.CourseCreateRequest) web.CourseResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	course := domain.Course{
		Name:           request.Name,
		Description:    request.Description,
		Price:          request.Price,
		Image:          request.Image,
		Video:          request.Video,
		InstructorName: request.InstructorName,
	}
	course = service.AdminRepository.CreateCourse(ctx, tx, course)

	return helper.ToCourseResponse(course)

}

func (service *AdminServiceImpl) UpdateCourse(ctx context.Context, request web.CourseUpdateRequest) web.CourseResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	course := domain.Course{
		Name:        request.Name,
		Description: request.Description,
	}
	course = service.AdminRepository.UpdateCourse(ctx, tx, course)

	return helper.ToCourseResponse(course)
}

func (service *AdminServiceImpl) DeleteCourse(ctx context.Context, courseId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.AdminRepository.DeleteCourse(ctx, tx, courseId)

}
