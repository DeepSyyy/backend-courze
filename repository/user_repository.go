package repository

import (
	"context"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	UpdateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
