package repository

import (
	"context"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type AdminRepository interface {
	CreateCourse(ctx context.Context, tx *sql.Tx, course domain.Course) domain.Course
	UpdateCourse(ctx context.Context, tx *sql.Tx, course domain.Course) domain.Course
	DeleteCourse(ctx context.Context, tx *sql.Tx, courseId int) domain.Course
}
