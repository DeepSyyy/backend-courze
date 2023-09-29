package controller

import (
	"courze-backend-app/helper"
	web "courze-backend-app/model/web"
	"courze-backend-app/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type InstructorControllerImpl struct {
	InstructorService service.InstructorService
}

func NewInstructorController(instructorService service.InstructorService) InstructorController {
	return &InstructorControllerImpl{
		InstructorService: instructorService,
	}
}

func (controller *InstructorControllerImpl) Insert(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	instructorCreateReq := web.InstructorCreateRequest{}
	helper.ReadFromRequestBody(request, &instructorCreateReq)

	instructorResponse := controller.InstructorService.Insert(request.Context(), instructorCreateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   instructorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
