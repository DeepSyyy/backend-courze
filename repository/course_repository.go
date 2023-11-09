package repository

import (
	"context"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type CourseRepository interface {
	GetAllCourse(ctx context.Context, tx *sql.Tx) []domain.Course
	GetCourseById(ctx context.Context, tx *sql.Tx, courseId int) (domain.Course, error)
	GetCourseByInstructorName(ctx context.Context, tx *sql.Tx, instructorName string) []domain.Course
}
