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
	if err != nil {
		// Tangani kesalahan parsing courseId
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest, // 400 Bad Request
			Status: "Bad Request",
			Data:   nil,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Memanggil GetCourseById dari CourseService
	courseResponse, err := controller.CourseService.GetCourseById(request.Context(), id)
	if err != nil {
		// Tangani kesalahan dari CourseService
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound, // 404 Not Found
			Status: "Not Found",
			Data:   err.Error(), // Menyertakan pesan kesalahan dalam respons
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Sukses mendapatkan course, berikan respons dengan data course
	webResponse := web.WebResponse{
		Code:   http.StatusOK, // 200 OK
		Status: "OK",
		Data:   courseResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CourseControllerImpl) GetCourseByInstructorId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	instructorName := params.ByName("instructorName")

	courseResponses := controller.CourseService.GetCourseByInstructorName(request.Context(), instructorName)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   courseResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
