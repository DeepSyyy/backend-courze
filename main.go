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
	// Course
	courseRepository := repository.NewCourseRepository()
	courseService := service.NewCourseService(courseRepository, db, validate)
	courseController := controller.NewCourseController(courseService)
	// Instructor
	instructorRepository := repository.NewInstructorRepository()
	instructorService := service.NewInstructorService(instructorRepository, db, validate)
	InstructorController := controller.NewInstructorController(instructorService)

	router := httprouter.New()

	router.POST("/api/courze/course", courseController.CreateCourse)
	router.POST("/api/courze/instructor", InstructorController.Insert)

	router.GET("/api/courze/course", courseController.GetAllCourse)
	router.GET("/api/courze/course/:courseId", courseController.GetCourseById)

	router.GET("/api/courze/instructor/:instructorId", courseController.GetCourseByInstructorId)
	address := "localhost:8080"
	fmt.Printf("server running on http://%v \n", address)
	// Menjalankan server HTTP dengan router yang telah Anda buat
	if err := http.ListenAndServe(address, router); err != nil {
		panic(err)
	}

}
