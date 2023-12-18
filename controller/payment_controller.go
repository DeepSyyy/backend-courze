package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PaymentController interface {
	CreatePayment(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetPaymentByUserID(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
