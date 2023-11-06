package repository

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
	"database/sql"

	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	user.Id = uuid.New()
	SQL := "INSERT INTO user(user_id,nama, email, password) VALUES(?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Email, user.Password)
	helper.PanicIfError(err)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) UpdateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE user SET name = ?, email = ?, password = ? WHERE user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}
