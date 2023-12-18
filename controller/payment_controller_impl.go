package controller

import (
	"courze-backend-app/helper"
	"courze-backend-app/model/web"
	service "courze-backend-app/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PaymentControllerImpl struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

func (controller *PaymentControllerImpl) CreatePayment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentCreateRequest := web.PaymentRequest{}
	helper.ReadFromRequestBody(request, &paymentCreateRequest)

	courseResponse := controller.PaymentService.CreatePayment(request.Context(), paymentCreateRequest)
	web := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   courseResponse,
	}

	helper.WriteToResponseBody(writer, web)
}

func (controller *PaymentControllerImpl) GetPaymentByUserID(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userID := param.ByName("userID")

	paymentResponses := controller.PaymentService.GetPaymentByUserID(request.Context(), userID)
	if paymentResponses == nil {
		// Tangani kesalahan dari CourseService
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   nil,
		}
		helper.WriteNotFoundToResponseBody(writer, webResponse)

	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   paymentResponses,
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}
