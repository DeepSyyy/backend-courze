package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	UpdateUser(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	LoginUser(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Enroll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	GetUserCourseByID(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	AddWishlist(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	GetWishlistByID(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
