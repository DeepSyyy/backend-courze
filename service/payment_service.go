package service

import (
	"context"
	"courze-backend-app/model/web"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, request web.PaymentRequest) web.PaymentResponse
	GetPaymentByUserID(ctx context.Context, userID string) []web.PaymentResponse
}
