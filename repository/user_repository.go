package repository

import (
	"context"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	UpdateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	LoginUser(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	GetUserByID(ctx context.Context, tx *sql.Tx, userID string) (domain.User, error)
}
