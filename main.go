package main

import (
	"todosAPI/routes"
	"todosAPI/src/auth"
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
	authRoute := routes.AuthRouteInit(auth.InitConstructorController(auth.InitConstructor()))

	v1 := router.Group("/api/v1")	

	/**
	 * register login route
	 * * register login route
	 */
	v1.POST("/registration",authRoute.RegistrationUserRoute)

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
	v1.DELETE("/todo/:id",todoRoutes.DeleteTodo)

	
	router.Run("localhost:8080")
}