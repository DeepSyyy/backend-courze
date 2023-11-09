package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	UpdateUser(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	LoginUser(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
