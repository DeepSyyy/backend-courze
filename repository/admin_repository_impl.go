package repository

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
	"database/sql"
	"time"
)

type AdminRepositoryImpl struct {
}

func NewAdminRepository() *AdminRepositoryImpl {
	return &AdminRepositoryImpl{}
}

func (repository *AdminRepositoryImpl) CreateCourse(ctx context.Context, tx *sql.Tx, course domain.Course) domain.Course {
	SQL := "INSERT INTO course(course_name, course_description, course_price, course_image, course_video, instructor_name, created_at, updated_at, sneakpeek) VALUES(?,?,?,?,?,?,?,?, (SELECT JSON_ARRAY(?)))"
	timeNow := time.Now()
	course.CreatedAt = timeNow
	course.UpdatedAt = timeNow
	result, err := tx.ExecContext(ctx, SQL, course.Name, course.Description, course.Price, course.Image, course.Video, course.InstructorName, course.CreatedAt, course.UpdatedAt, course.SneakPeak)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	course.Id = int(id)

	return course
}

func (repository *AdminRepositoryImpl) UpdateCourse(ctx context.Context, tx *sql.Tx, course domain.Course) domain.Course {
	SQL := "UPDATE course SET course_name = ?, course_description = ?, course_price = ?, course_image = ?, course_video = ?, instructor_id = ?, sneakpeek = (SELECT JSON_ARRAY(?)) WHERE course_id = ?"
	_, err := tx.ExecContext(ctx, SQL, course.Name, course.Description, course.Price, course.Image, course.Video, course.InstructorName, course.SneakPeak, course.Id)
	helper.PanicIfError(err)

	return course
}

func (repository *AdminRepositoryImpl) DeleteCourse(ctx context.Context, tx *sql.Tx, courseId int) domain.Course {
	SQL := "DELETE FROM course WHERE course_id = ?"
	_, err := tx.ExecContext(ctx, SQL, courseId)
	helper.PanicIfError(err)

	course := domain.Course{}
	course.Id = courseId

	return course
}
