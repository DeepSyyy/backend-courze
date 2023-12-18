package repository

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (repository *PaymentRepositoryImpl) CreatePayment(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment {
	SQL := "INSERT INTO payment(id,user_id, course_id, price, method, created_at) VALUES(?,?,?,?,?,?))"
	timeNow := time.Now()
	payment.CreatedAt = timeNow
	payment.Id = uuid.New()
	_, err := tx.ExecContext(ctx, SQL, payment.Id, payment.UserId, payment.CourseId, payment.Price, payment.Method, payment.CreatedAt)
	helper.PanicIfError(err)

	return payment
}

func (repository *PaymentRepositoryImpl) GetPaymentByUserID(ctx context.Context, tx *sql.Tx, userID string) ([]domain.Payment, error) {
	SQL := "SELECT id,user_id,course_id, price, method, created_at FROM payment WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userID)
	helper.PanicIfError(err)
	defer rows.Close()

	var result []domain.Payment
	for rows.Next() {
		payment := domain.Payment{}
		err := rows.Scan(&payment.Id, &payment.UserId, &payment.CourseId, &payment.Price, &payment.Method, &payment.CreatedAt)
		helper.PanicIfError(err)
		result = append(result, payment)
	}

	return result, nil
}
