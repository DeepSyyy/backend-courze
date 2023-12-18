package repository

import (
	"context"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment
	GetPaymentByUserID(ctx context.Context, tx *sql.Tx, userID string) ([]domain.Payment, error)
}
