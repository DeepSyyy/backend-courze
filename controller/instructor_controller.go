package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type InstructorController interface {
	Insert(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
