package main

import (
	"github.com/Ivanrizkys/go-api/api/config"
	"github.com/Ivanrizkys/go-api/api/controller"
	"github.com/Ivanrizkys/go-api/api/repository"
	"github.com/Ivanrizkys/go-api/api/router"
	"github.com/Ivanrizkys/go-api/api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := config.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// router := gin.Default()

	// router.GET("/api/categories", categoryController.FindAll)

	router := router.NewRouter(categoryController)

	router.Run()
}
