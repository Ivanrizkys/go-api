package controller

import (
	"github.com/Ivanrizkys/go-api/api/service"
	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) FindAll(c *gin.Context) {
	categoryResponse := controller.CategoryService.FindAll(c.Request.Context())
	c.JSON(200, gin.H{
		"code":   200,
		"status": "OK",
		"data":   categoryResponse,
	})
}
