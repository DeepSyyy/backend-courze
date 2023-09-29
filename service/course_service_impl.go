package service

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
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

func (service *CourseServiceImpl) CreateCourse(ctx context.Context, request web.CourseCreateRequest) web.CourseResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	course := domain.Course{
		CourseName:        request.Name,
		CourseDescription: request.Description,
		CoursePrice:       request.Price,
		CourseImage:       request.Image,
		CourseVideo:       request.Video,
		InstructorId:      request.InstructorId,
	}

	course = service.CourseRepository.CreateCourse(ctx, tx, course)
	return helper.ToCourseResponse(course)
}

func (service *CourseServiceImpl) GetAllCourse(ctx context.Context) []web.CourseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	courses := service.CourseRepository.GetAllCourse(ctx, tx)
	return helper.ToCourseResponses(courses)
}

func (service *CourseServiceImpl) GetCourseById(ctx context.Context, courseId int) web.CourseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	course := service.CourseRepository.GetCourseById(ctx, tx, courseId)
	return helper.ToCourseResponse(course)
}

func (service *CourseServiceImpl) GetCourseByInstructorId(ctx context.Context, instructorId int) []web.CourseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	courses := service.CourseRepository.GetCourseByInstructorId(ctx, tx, instructorId)
	return helper.ToCourseResponses(courses)
}
