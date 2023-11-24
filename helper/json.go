package helper

import (
	web "courze-backend-app/model/web"
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(response http.ResponseWriter, result interface{}) {
	response.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(response)
	err := encoder.Encode(result)
	PanicIfError(err)
}

func ToErrorResponse(err error, statusCode int) web.WebResponse {
	return web.WebResponse{
		Code:   statusCode,
		Status: "Error",
		Data:   err.Error(),
	}
}
