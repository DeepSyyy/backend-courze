package repository

import (
	"context"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type CourseRepository interface {
	CreateCourse(ctx context.Context, tx *sql.Tx, course domain.Course) domain.Course
	GetAllCourse(ctx context.Context, tx *sql.Tx) []domain.Course
	GetCourseById(ctx context.Context, tx *sql.Tx, courseId int) domain.Course
	GetCourseByInstructorId(ctx context.Context, tx *sql.Tx, instructorId int) []domain.Course
}
