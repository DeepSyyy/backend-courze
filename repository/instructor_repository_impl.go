package repository

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type InstructorRepositoryImpl struct {
}

func NewInstructorRepository() InstructorRepository {
	return &InstructorRepositoryImpl{}
}

func (repository *InstructorRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, instructor domain.Instructor) domain.Instructor {
	SQL := "INSERT INTO instructor (name, email) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, instructor.Name, instructor.Email)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	instructor.Id = int(id)
	return instructor
}
