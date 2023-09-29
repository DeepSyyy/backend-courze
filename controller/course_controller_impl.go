package controller

import (
	"courze-backend-app/helper"
	web "courze-backend-app/model/web"
	service "courze-backend-app/service"
	"net/http"
	"strconv"

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

func (controller *CourseControllerImpl) GetCourseById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	courseId := params.ByName("courseId")
	id, err := strconv.Atoi(courseId)
	helper.PanicIfError(err)

	courseResponse := controller.CourseService.GetCourseById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   courseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CourseControllerImpl) GetCourseByInstructorId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	instructorId := params.ByName("instructorId")
	id, err := strconv.Atoi(instructorId)
	helper.PanicIfError(err)

	courseResponses := controller.CourseService.GetCourseByInstructorId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   courseResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
