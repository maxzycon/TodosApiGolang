package main

import (
	"todosAPI/category"
	"todosAPI/routes"
	"todosAPI/todos"
	"todosAPI/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.ConnectDB()
	router := gin.Default()

	todoRoutes := routes.TodoRouteInit(todos.InitConstructorController(todos.InitConstructor()))
	categoryRoutes := routes.CategoryRouteInit(category.InitConstructorController(category.InitConstructor()))

	v1 := router.Group("/api/v1")	
	
	// category
	v1.GET("/category",categoryRoutes.GetCategoryRoute)
	v1.GET("/category/:id",categoryRoutes.FindCategoryRoute)
	v1.POST("/category",categoryRoutes.CreateCategoryRoute)
	v1.DELETE("/category/:id",categoryRoutes.DeleteCategoryRoute)
	// end category

	v1.GET("/todo",todoRoutes.ReadTodo)
	v1.POST("/todo",todoRoutes.CreateTodo)
	router.Run("localhost:8080")
}