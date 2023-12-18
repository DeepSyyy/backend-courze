package service

import (
	"context"
	"courze-backend-app/model/domain"
	"courze-backend-app/model/web"
	repository "courze-backend-app/repository"
	"database/sql"

	"courze-backend-app/helper"

	"github.com/go-playground/validator"
)

type PaymentServiceImpl struct {
	PaymentRepository repository.PaymentRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewPaymentService(paymentRepository repository.PaymentRepository, DB *sql.DB, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *PaymentServiceImpl) CreatePayment(ctx context.Context, request web.PaymentRequest) web.PaymentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payment := domain.Payment{
		UserId:    request.UserId,
		CourseId:  request.CourseId,
		Price:     request.Price,
		Method:    request.Method,
		CreatedAt: helper.GetCurrentTime(),
	}
	payment = service.PaymentRepository.CreatePayment(ctx, tx, payment)

	return helper.ToPaymentResponse(payment)
}

func (service *PaymentServiceImpl) GetPaymentByUserID(ctx context.Context, userID string) []web.PaymentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payments, _ := service.PaymentRepository.GetPaymentByUserID(ctx, tx, userID)
	return helper.ToPaymentResponses(payments)
}
