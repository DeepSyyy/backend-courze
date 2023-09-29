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
	courseRepository := repository.NewCourseRepository()
	courseService := service.NewCourseService(courseRepository, db, validate)
	courseController := controller.NewCourseController(courseService)

	router := httprouter.New()

	router.POST("/api/courses", courseController.CreateCourse)
	router.GET("/api/courses", courseController.GetAllCourse)
	address := "localhost:8080"
	fmt.Printf("server running on http://%v", address)
	// Menjalankan server HTTP dengan router yang telah Anda buat
	if err := http.ListenAndServe(address, router); err != nil {
		panic(err)
	}

}
