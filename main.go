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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	utils.ConnectDB()
	router := gin.Default()
	router.Use(CORSMiddleware())

	/**
	 * init
	 * * initial route handler
	 */
	todoRoutes := routes.TodoRouteInit(todos.InitConstructorController(todos.InitConstructor()))
	categoryRoutes := routes.CategoryRouteInit(category.InitConstructorController(category.InitConstructor()))
	authRoute := routes.AuthRouteInit(auth.InitConstructorController(auth.InitConstructor()))

	// * route group v1 e.g domain.com/api/v1
	v1 := router.Group("/api/v1")
	token, _ := token.NewJWTMaker("12345678901234567890123456789012")
	middlewareAuth := v1.Group("/").Use(middleware.AuthMiddleware(token))

	/**
	 * register login route
	 * * register login route
	 */
	v1.POST("/registration", authRoute.RegistrationUserRoute)
	v1.POST("/login", authRoute.LoginUserRoute)

	/**
	 * category route
	 * * this is category handler route
	 */
	v1.GET("/category", categoryRoutes.GetCategoryRoute)
	// middlewareAuth.GET("/category",categoryRoutes.GetCategoryRoute)
	middlewareAuth.GET("/category/:id", categoryRoutes.FindCategoryRoute)
	middlewareAuth.POST("/category", categoryRoutes.CreateCategoryRoute)
	middlewareAuth.DELETE("/category/:id", categoryRoutes.DeleteCategoryRoute)

	/**
	 * _todo route
	 * * this is todo handler route
	 */
	middlewareAuth.GET("/todo", todoRoutes.ReadTodo)
	middlewareAuth.POST("/todo", todoRoutes.CreateTodo)
	middlewareAuth.PUT("/todo/:id", todoRoutes.UpdateTodo)
	middlewareAuth.DELETE("/todo/:id", todoRoutes.DeleteTodo)

	router.Run("0.0.0.0:8080")
}
