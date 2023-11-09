package repository

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	user.Id = uuid.New()
	user.IsVerified = false
	user.CreatedAt = helper.GetCurrentTime()
	SQL := "INSERT INTO user(user_id,nama, email, password, is_verified, created_at) VALUES(?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Email, user.Password, user.IsVerified, user.CreatedAt)
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

func (repository *UserRepositoryImpl) LoginUser(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "SELECT email, password FROM user WHERE email = ?"
	row := tx.QueryRowContext(ctx, SQL, user.Email)

	var result domain.User
	err := row.Scan(&result.Email, &result.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Email tidak ditemukan
			return domain.User{}, errors.New("email not found")
		}

		// Kesalahan lainnya
		return domain.User{}, err
	}

	// Periksa apakah password cocok (dengan asumsi password di-hashing)
	if result.Password != user.Password {
		// Password tidak cocok
		return domain.User{}, errors.New("invalid password")
	}

	return result, nil
}

func (repository *UserRepositoryImpl) GetUserByID(ctx context.Context, tx *sql.Tx, userID string) (domain.User, error) {
	SQL := "SELECT user_id, name, email, password, is_verified, created_at FROM user WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userID)
	helper.PanicIfError(err)
	defer rows.Close()

	var result domain.User
	if rows.Next() {
		err := rows.Scan(&result.Id, &result.Name, &result.Email, &result.Password, &result.IsVerified, &result.CreatedAt)
		helper.PanicIfError(err)
		return result, nil
	} else {
		// User tidak ditemukan, mungkin berikan respons atau tindakan yang sesuai
		return result, errors.New("user tidak ditemukan")
	}
}
