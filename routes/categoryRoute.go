package routes

import (
	"errors"
	"net/http"
	"todosAPI/src/category"
	"todosAPI/src/category/create"
	"todosAPI/src/category/read"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoryRoute struct {
	categoryCreate category.Controller
}

func CategoryRouteInit(todoController category.Controller) *categoryRoute  {
	return &categoryRoute{todoController}
}

func (r *categoryRoute) GetCategoryRoute(c *gin.Context) {
	category,err_response := r.categoryCreate.GetCategoryController()
	
	if err_response != nil {
		c.JSON(http.StatusUnprocessableEntity,err_response)
		return
	}
	
	formatResponse := read.FormatCategories(category)
	c.JSON(http.StatusOK,formatResponse)
}

func (r *categoryRoute) FindCategoryRoute(c *gin.Context) {
	id := c.Param("id")
	category,err := r.categoryCreate.FindCategoryController(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound,gin.H{"status":"id not found"})
		return
	}
	
	formatResponse := create.FormatCategory(category)
	c.JSON(http.StatusOK,formatResponse)
}

func (r *categoryRoute) DeleteCategoryRoute(c *gin.Context) {
	id := c.Param("id")
	_,err := r.categoryCreate.DeleteCategoryController(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound,gin.H{"status":"id not found"})
		return
	}
	
	c.JSON(http.StatusOK,gin.H{"status":"success delete"})
}

func (r *categoryRoute) CreateCategoryRoute(c *gin.Context) {
	var input create.CreateCategoryInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}

	category,err_response := r.categoryCreate.CreateCategoryController(input)
	
	if err_response != nil {
		c.JSON(http.StatusUnprocessableEntity,err_response)
		return
	}

	formatResponse := create.FormatCategory(category)
	c.JSON(http.StatusOK,formatResponse)
}