package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AdminController interface {
	CreateCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteCourse(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
