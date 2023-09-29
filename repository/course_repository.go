package repository

import (
	"context"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type CourseRepository interface {
	CreateCourse(ctx context.Context, tx *sql.Tx, course domain.Course) domain.Course
	GetAllCourse(ctx context.Context, tx *sql.Tx) []domain.Course
}
