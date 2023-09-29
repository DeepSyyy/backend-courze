package repository

import (
	"context"
	"courze-backend-app/model/domain"
	"database/sql"
)

type InstructorRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, instructor domain.Instructor) domain.Instructor
}
