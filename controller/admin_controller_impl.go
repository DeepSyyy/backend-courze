package controller

import (
	"courze-backend-app/helper"
	"courze-backend-app/model/web"
	service "courze-backend-app/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AdminControllerImpl struct {
	AdminService service.AdminService
}

func NewAdminController(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		AdminService: adminService,
	}
}

func (controller *AdminControllerImpl) CreateCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	courseCreateRequest := web.CourseCreateRequest{}
	helper.ReadFromRequestBody(request, &courseCreateRequest)

	courseResponse := controller.AdminService.CreateCourse(request.Context(), courseCreateRequest)
	web := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   courseResponse,
	}

	helper.WriteToResponseBody(writer, web)
}

func (controller *AdminControllerImpl) UpdateCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	courseUpdateRequest := web.CourseUpdateRequest{}
	helper.ReadFromRequestBody(request, &courseUpdateRequest)

	courseResponse := controller.AdminService.UpdateCourse(request.Context(), courseUpdateRequest)
	web := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   courseResponse,
	}

	helper.WriteToResponseBody(writer, web)
}

func (controller *AdminControllerImpl) DeleteCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	courseId := params.ByName("courseId")
	id, err := strconv.Atoi(courseId)
	helper.PanicIfError(err)

	controller.AdminService.DeleteCourse(request.Context(), id)
	web := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, web)
}
