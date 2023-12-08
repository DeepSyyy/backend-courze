package helper

import (
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

func WriteNotFoundToResponseBody(response http.ResponseWriter, result interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusNotFound)
	encoder := json.NewEncoder(response)
	err := encoder.Encode(result)
	PanicIfError(err)
}

func WriteBadRequestToResponseBody(response http.ResponseWriter, result interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)
	encoder := json.NewEncoder(response)
	err := encoder.Encode(result)
	PanicIfError(err)
}

func WriteInternalServerErrorToResponseBody(response http.ResponseWriter, result interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	encoder := json.NewEncoder(response)
	err := encoder.Encode(result)
	PanicIfError(err)
}

func WriteSuccesCreatedToResponseBody(response http.ResponseWriter, result interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(response)
	err := encoder.Encode(result)
	PanicIfError(err)
}

func WriteDeleteSuccessToResponseBody(response http.ResponseWriter, result interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(response)
	err := encoder.Encode(result)
	PanicIfError(err)
}

func WriteUpdateSuccessToResponseBody(response http.ResponseWriter, result interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(response)
	err := encoder.Encode(result)
	PanicIfError(err)
}
