package router

import (
	"github.com/Ivanrizkys/go-api/api/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(categoryController controller.CategoryController) *gin.Engine {
	router := gin.Default()

	router.GET("/api/categories", categoryController.FindAll)

	return router
}
