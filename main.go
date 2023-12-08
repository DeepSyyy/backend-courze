package main

import (
	"courze-backend-app/app"
	"courze-backend-app/controller"
	repository "courze-backend-app/repository"
	service "courze-backend-app/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	//repository
	// Course
	courseRepository := repository.NewCourseRepository()
	courseService := service.NewCourseService(courseRepository, db, validate)
	courseController := controller.NewCourseController(courseService)
	//admin
	adminRepository := repository.NewAdminRepository()
	adminService := service.NewAdminService(adminRepository, db, validate)
	adminController := controller.NewAdminController(adminService)
	//user
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()
	//admin
	router.POST("/api/courze/admin", adminController.CreateCourse)
	//course
	router.GET("/api/courze/course", courseController.GetAllCourse)
	router.GET("/api/courze/course/name/:courseName", courseController.GetCourseByName)
	router.GET("/api/courze/course/id/:courseId", courseController.GetCourseById)
	router.GET("/api/courze/instructor/:instructorId", courseController.GetCourseByInstructorId)
	//user
	router.POST("/api/courze/user/register", userController.Register)
	router.POST("/api/courze/user/login", userController.LoginUser)
	router.PUT("/api/courze/user", userController.UpdateUser)
	address := "localhost:8080"
	fmt.Printf("server running on http://%v \n", address)
	// Menjalankan server HTTP dengan router yang telah Anda buat
	if err := http.ListenAndServe(address, router); err != nil {
		panic(err)
	}

}
