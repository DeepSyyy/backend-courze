package controller

import (
	"courze-backend-app/helper"
	"courze-backend-app/model/domain"
	"courze-backend-app/model/web"
	"courze-backend-app/service"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userCreateReq := web.UserRequest{}
	helper.ReadFromRequestBody(request, &userCreateReq)

	userResponse, err := controller.UserService.Register(request.Context(), userCreateReq)

	if err != nil {
		switch err := err.(type) {
		case *domain.ErrUserAlreadyExists:
			// Handle ErrUserAlreadyExists
			fmt.Printf("Type: %T\n", err)
			webResponse := web.WebResponse{
				Code:   http.StatusConflict,
				Status: "Conflict",
				Data:   "User already exists",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		default:
			// Handle other error cases
			fmt.Printf("Type: %T\n", err)
			webResponse := web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Error",
				Data:   "An unexpected error occurred",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}

	}

	// Success
	helper.WriteToResponseBody(writer, userResponse)
}

func (controller *UserControllerImpl) UpdateUser(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userUpdateReq := web.UserUpdateRequest{}
	userID, err := uuid.Parse(param.ByName("userId"))
	if err != nil {
		// Tangani kesalahan parsing UUID
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalid user ID",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Memanggil GetUserByID dengan userID yang benar
	_, err = controller.UserService.GetUserByID(request.Context(), userID.String())
	if err != nil {
		// Tangani kesalahan pengambilan pengguna
		webResponse := web.WebResponse{
			Code:   404,
			Status: "Not Found",
			Data:   "User not found",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	helper.ReadFromRequestBody(request, &userUpdateReq)

	userResponse := controller.UserService.UpdateUser(request.Context(), userUpdateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) LoginUser(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userLoginReq := web.UserRequest{}
	helper.ReadFromRequestBody(request, &userLoginReq)

	// Memanggil LoginUser dari UserService
	userResponse, err := controller.UserService.LoginUser(request.Context(), userLoginReq)
	if err != nil {
		// Tangani kesalahan login dengan memberikan respons yang sesuai
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound, // 404 Not Found
			Status: "Not Found",
			Data:   err.Error(), // Menyertakan pesan kesalahan dalam respons
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Login sukses, berikan respons dengan data pengguna yang masuk
	webResponse := web.WebResponse{
		Code:   http.StatusOK, // 200 OK
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Enroll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userEnrollReq := web.UserCourseRequest{}
	helper.ReadFromRequestBody(request, &userEnrollReq)

	// Memanggil LoginUser dari UserService
	userResponse, err := controller.UserService.Enroll(request.Context(), userEnrollReq)
	if err != nil {
		// Tangani kesalahan login dengan memberikan respons yang sesuai
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound, // 404 Not Found
			Status: "Not Found",
			Data:   err.Error(), // Menyertakan pesan kesalahan dalam respons
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Login sukses, berikan respons dengan data pengguna yang masuk
	webResponse := web.WebResponse{
		Code:   http.StatusOK, // 200 OK
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) GetUserCourseByID(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userID := param.ByName("userID")

	userCourseResponses := controller.UserService.GetUserCourseByID(request.Context(), userID)
	if userCourseResponses == nil {
		// Tangani kesalahan dari CourseService
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   nil,
		}
		helper.WriteNotFoundToResponseBody(writer, webResponse)

	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   userCourseResponses,
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}

func (controller *UserControllerImpl) AddWishlist(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userWishReq := web.WishlistRequest{}
	helper.ReadFromRequestBody(request, &userWishReq)

	// Memanggil LoginUser dari UserService
	userResponse, err := controller.UserService.AddWishlist(request.Context(), userWishReq)
	if err != nil {
		// Tangani kesalahan login dengan memberikan respons yang sesuai
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound, // 404 Not Found
			Status: "Not Found",
			Data:   err.Error(), // Menyertakan pesan kesalahan dalam respons
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Login sukses, berikan respons dengan data pengguna yang masuk
	webResponse := web.WebResponse{
		Code:   http.StatusOK, // 200 OK
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) GetWishlistByID(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userID := param.ByName("userID")

	wishlistResponses := controller.UserService.GetWishlistByID(request.Context(), userID)
	if wishlistResponses == nil {
		// Tangani kesalahan dari CourseService
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   nil,
		}
		helper.WriteNotFoundToResponseBody(writer, webResponse)

	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   wishlistResponses,
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}
