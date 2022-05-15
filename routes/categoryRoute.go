package routes

import (
	"errors"
	"net/http"
	"todosAPI/src/category"
	"todosAPI/src/category/create"
	"todosAPI/src/category/read"
	"todosAPI/src/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// * bikin sebuah data yang mana data typenya adalah controller
type categoryRoute struct {
	categoryController category.Controller
}

// * masukan category ke categoryRouteStruct
func CategoryRouteInit(todoController category.Controller) *categoryRoute {
	return &categoryRoute{todoController}
}

// * declare categoryRoute yg mana dalamnya sudah ada controller yang terisi
func (r *categoryRoute) GetCategoryRoute(c *gin.Context) {
	todoQuery := c.Query("todo")
	// * memanggil controller yang terlah diisi valuenya saat CategoryRouteInit
	category, err_response := r.categoryController.GetCategoryController(todoQuery)

	if err_response != nil {
		c.JSON(http.StatusUnprocessableEntity, err_response)
		return
	}

	if todoQuery == "1" {
		formatResponse := read.FormatCategories(category)
		c.JSON(http.StatusOK, formatResponse)
	} else {
		formatResponse := read.FormatWithoutTodosCategories(category)
		c.JSON(http.StatusOK, formatResponse)
	}
}

func (r *categoryRoute) FindCategoryRoute(c *gin.Context) {
	id := c.Param("id")
	category, err := r.categoryController.FindCategoryController(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"status": "id not found"})
		return
	}

	formatResponse := create.FormatCategory(category)
	c.JSON(http.StatusOK, formatResponse)
}

func (r *categoryRoute) DeleteCategoryRoute(c *gin.Context) {
	id := c.Param("id")
	_, err := r.categoryController.DeleteCategoryController(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"status": "id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success delete"})
}

func (r *categoryRoute) CreateCategoryRoute(c *gin.Context) {
	// var ctx *gin.Context
	authPayload := c.MustGet("authorization_payload_key").(*token.Payload)

	var input create.CreateCategoryInput
	err := c.ShouldBindJSON(&input)
	input.UserID = authPayload.Username

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	category, err_response := r.categoryController.CreateCategoryController(input)

	if err_response != nil {
		c.JSON(http.StatusUnprocessableEntity, err_response)
		return
	}

	formatResponse := create.FormatCategory(category)
	c.JSON(http.StatusOK, formatResponse)
}
