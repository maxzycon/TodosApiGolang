package main

import (
	"todosAPI/middleware"
	"todosAPI/routes"
	"todosAPI/src/auth"
	"todosAPI/src/category"
	"todosAPI/src/todos"
	"todosAPI/src/token"
	"todosAPI/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.ConnectDB()
	router := gin.Default()

	/**
	 * init
	 * * initial route handler
	 */
	todoRoutes := routes.TodoRouteInit(todos.InitConstructorController(todos.InitConstructor()))
	categoryRoutes := routes.CategoryRouteInit(category.InitConstructorController(category.InitConstructor()))
	authRoute := routes.AuthRouteInit(auth.InitConstructorController(auth.InitConstructor()))

	// * route group v1 e.g domain.com/api/v1
	v1 := router.Group("/api/v1")	
	token,_ := token.NewJWTMaker("12345678901234567890123456789012")
	middlewareAuth := v1.Group("/").Use(middleware.AuthMiddleware(token))

	/**
	 * register login route
	 * * register login route
	 */
	v1.POST("/registration",authRoute.RegistrationUserRoute)
	v1.POST("/login",authRoute.LoginUserRoute)

	/**
	 * category route
	 * * this is category handler route 
	 */
	middlewareAuth.GET("/category",categoryRoutes.GetCategoryRoute)
	middlewareAuth.GET("/category/:id",categoryRoutes.FindCategoryRoute)
	middlewareAuth.POST("/category",categoryRoutes.CreateCategoryRoute)
	middlewareAuth.DELETE("/category/:id",categoryRoutes.DeleteCategoryRoute)

	/**
	 * _todo route
	 * * this is todo handler route
	 */
	middlewareAuth.GET("/todo",todoRoutes.ReadTodo)
	middlewareAuth.POST("/todo",todoRoutes.CreateTodo)
	middlewareAuth.PUT("/todo/:id",todoRoutes.UpdateTodo)
	middlewareAuth.DELETE("/todo/:id",todoRoutes.DeleteTodo)

	
	router.Run("localhost:8080")
}