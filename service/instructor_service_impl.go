package service

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
	web "courze-backend-app/model/web"
	"courze-backend-app/repository"
	"database/sql"

	"github.com/go-playground/validator"
)

type InstructorServiceImpl struct {
	InstructorRepository repository.InstructorRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewInstructorService(instructorRepository repository.InstructorRepository, DB *sql.DB, Validate *validator.Validate) InstructorService {
	return &InstructorServiceImpl{
		InstructorRepository: instructorRepository,
		DB:                   DB,
		Validate:             Validate,
	}
}

func (service *InstructorServiceImpl) Insert(ctx context.Context, request web.InstructorCreateRequest) web.InstructorResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	instructor := domain.Instructor{
		Name:  request.Name,
		Email: request.Email,
	}

	instructor = service.InstructorRepository.Insert(ctx, tx, instructor)
	return helper.ToInstructorResponse(instructor)
}
