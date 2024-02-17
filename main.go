package main

import (
	"4kbr/restful-golang/app"
	"4kbr/restful-golang/controller"
	"4kbr/restful-golang/exception"
	"4kbr/restful-golang/helper"
	"4kbr/restful-golang/middleware"
	"4kbr/restful-golang/repository"
	"4kbr/restful-golang/service"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Run on http://localhost:" + port)
	server := http.Server{
		Addr:    `localhost:` + port,
		Handler: middleware.NewAuthMiddleWare(router),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Run on http://localhost:" + port)
}
