package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CourseController interface {
	CreateCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAllCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetCourseById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetCourseByInstructorId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
