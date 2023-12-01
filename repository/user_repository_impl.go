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

func (repository *UserRepositoryImpl) GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	// Prepare SQL query
	SQL := "SELECT user_id, nama, email, password, is_verified FROM user WHERE email = ?"

	// Execute SQL query
	row := tx.QueryRowContext(ctx, SQL, email)

	// Scan results into user object
	var user domain.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.IsVerified)

	// Return user object or error
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, nil
		}

		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	// Check if user with the given email already exists
	existingUser, err := repository.GetUserByEmail(ctx, tx, user.Email)
	if err != nil && err != sql.ErrNoRows {
		return domain.User{}, err
	}
	if existingUser != (domain.User{}) {
		return domain.User{}, errors.New("user already exists")
	}

	// Proceed with user registration if user doesn't exist
	user.Id = uuid.New()
	user.IsVerified = false
	user.CreatedAt = helper.GetCurrentTime()

	SQL := "INSERT INTO user(user_id, nama, email, password, is_verified, created_at) VALUES(?,?,?,?,?,?)"
	_, err = tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Email, user.Password, user.IsVerified, user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	// Return the registered user
	return user, nil
}
func (repository *UserRepositoryImpl) UpdateUser(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE user SET name = ?, email = ?, password = ? WHERE user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) LoginUser(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "SELECT user_id, nama, email, password,is_verified FROM user WHERE email = ?"
	row := tx.QueryRowContext(ctx, SQL, user.Email)

	var result domain.User
	err := row.Scan(&result.Id, &result.Name, &result.Email, &result.Password, &result.IsVerified)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("invalid email or password")
		}

		return domain.User{}, err
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

func (repository *UserRepositoryImpl) Enroll(ctx context.Context, tx *sql.Tx, usercourse domain.UserCourse) (domain.UserCourse, error) {
	SQL := "INSERT INTO user_courses(user_id,course_id,) VALUES(?,?))"
	result, err := tx.ExecContext(ctx, SQL, usercourse.UserId, usercourse.CourseId)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	usercourse.Id = int(id)

	return usercourse, nil
}

func (repository *UserRepositoryImpl) GetUserCourseByID(ctx context.Context, tx *sql.Tx, userID string) ([]domain.UserCourse, error) {
	SQL := "SELECT id,user_id,course_id, FROM user_courses WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userID)
	helper.PanicIfError(err)
	defer rows.Close()

	var result []domain.UserCourse
	if rows.Next() {
		course := domain.UserCourse{}
		err := rows.Scan(&course.Id, &course.UserId, &course.CourseId)
		helper.PanicIfError(err)

		result = append(result, course)
		helper.PanicIfError(err)
		return result, nil
	} else {
		// User course tidak ditemukan, mungkin berikan respons atau tindakan yang sesuai
		return result, errors.New("user course tidak ditemukan")
	}
}
