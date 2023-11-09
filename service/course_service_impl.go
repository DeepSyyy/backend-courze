package service

import (
	"context"
	"courze-backend-app/helper"
	web "courze-backend-app/model/web"
	repository "courze-backend-app/repository"
	"database/sql"

	"github.com/go-playground/validator"
)

type CourseServiceImpl struct {
	CourseRepository repository.CourseRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewCourseService(courseRepository repository.CourseRepository, DB *sql.DB, validate *validator.Validate) CourseService {
	return &CourseServiceImpl{
		CourseRepository: courseRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *CourseServiceImpl) GetAllCourse(ctx context.Context) []web.CourseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	courses := service.CourseRepository.GetAllCourse(ctx, tx)
	return helper.ToCourseResponses(courses)
}

func (service *CourseServiceImpl) GetCourseById(ctx context.Context, courseId int) (web.CourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		// Tangani kesalahan pembukaan transaksi
		return web.CourseResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	course, err := service.CourseRepository.GetCourseById(ctx, tx, courseId)
	if err != nil {
		// Tangani kesalahan pengambilan course
		return web.CourseResponse{}, err
	}

	// Menutup transaksi setelah sukses mendapatkan course

	return helper.ToCourseResponse(course), nil
}

func (service *CourseServiceImpl) GetCourseByInstructorName(ctx context.Context, instructorName string) []web.CourseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	courses := service.CourseRepository.GetCourseByInstructorName(ctx, tx, instructorName)
	return helper.ToCourseResponses(courses)
}
