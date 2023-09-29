package controller

import (
	"courze-backend-app/helper"
	web "courze-backend-app/model/web"
	service "courze-backend-app/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CourseControllerImpl struct {
	CourseService service.CourseService
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &CourseControllerImpl{
		CourseService: courseService,
	}
}

func (controller *CourseControllerImpl) CreateCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	courseCreateReq := web.CourseCreateRequest{}
	helper.ReadFromRequestBody(request, &courseCreateReq)

	courseResponse := controller.CourseService.CreateCourse(request.Context(), courseCreateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   courseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CourseControllerImpl) GetAllCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	courseResponses := controller.CourseService.GetAllCourse(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   courseResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
