package main

import (
	"todosAPI/routes"
	"todosAPI/src/category"
	"todosAPI/src/todos"
	"todosAPI/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.ConnectDB()
	router := gin.Default()

	todoRoutes := routes.TodoRouteInit(todos.InitConstructorController(todos.InitConstructor()))
	categoryRoutes := routes.CategoryRouteInit(category.InitConstructorController(category.InitConstructor()))

	v1 := router.Group("/api/v1")	
	
	/**
	 * category route
	 * * this is category handler route 
	 */
	v1.GET("/category",categoryRoutes.GetCategoryRoute)
	v1.GET("/category/:id",categoryRoutes.FindCategoryRoute)
	v1.POST("/category",categoryRoutes.CreateCategoryRoute)
	v1.DELETE("/category/:id",categoryRoutes.DeleteCategoryRoute)

	/**
	 * _todo route
	 * * this is todo handler route
	 */
	v1.GET("/todo",todoRoutes.ReadTodo)
	v1.POST("/todo",todoRoutes.CreateTodo)
	v1.PUT("/todo/:id",todoRoutes.UpdateTodo)

	
	router.Run("localhost:8080")
}