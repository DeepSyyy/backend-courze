package helper

import (
	domain "courze-backend-app/model/domain"
	web "courze-backend-app/model/web"
)

func ToPaymentResponse(payment domain.Payment) web.PaymentResponse {
	return web.PaymentResponse{
		Id:        payment.Id,
		UserId:    payment.UserId,
		CourseId:  payment.CourseId,
		Price:     payment.Price,
		Method:    payment.Method,
		CreatedAt: payment.CreatedAt,
	}
}

func ToPaymentResponses(payments []domain.Payment) []web.PaymentResponse {
	var paymentResponse []web.PaymentResponse
	for _, payment := range payments {
		paymentResponse = append(paymentResponse, ToPaymentResponse(payment))
	}
	return paymentResponse
}
