package routes

import (
	"errors"
	"net/http"
	"todosAPI/src/todos"
	"todosAPI/src/todos/create"
	"todosAPI/src/todos/read"
	"todosAPI/src/todos/update"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type todoRoute struct {
	todoController todos.Controller
}

func TodoRouteInit(todoController todos.Controller) *todoRoute  {
	return &todoRoute{todoController}
}

/**
 * * Handler create todo
 */

func (r *todoRoute) CreateTodo(c *gin.Context) {
	var input create.CreateTodoInput
	err := c.ShouldBindJSON(&input)
	
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	todos,err_response := r.todoController.CreateTodo(input)
	
	if err_response != nil {
		c.JSON(http.StatusUnprocessableEntity,err_response)
		return
	}
	formatResponse := create.FormatTodo(todos)
	c.JSON(http.StatusOK,formatResponse)
}

/**
 * * Handler read todo
 */

func (r *todoRoute) ReadTodo(c *gin.Context) {
	todos,err := r.todoController.ReadTodo()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity,err)
		return
	}
	formatResponse := read.FormatTodos(todos)
	c.JSON(http.StatusOK,formatResponse)
}

func (r *todoRoute) UpdateTodo(c *gin.Context) { 
	var input update.UpdateTodo
	/**
	 * * Get params id from uri
	 */
	id := c.Param("id")

	/**
	 * binding json input
	 * * validate json request
	 */
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	/**
	 * update service
	 * * passing input to service
	 */
	todo,err := r.todoController.UpdateTodo(input,id)

	/**
	 * checking has error service
	 * * check if service error or not
	 */
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound,gin.H{"status": "data not found"})
		}
		return
	}

	/**
	 * formatter
	 * * format the response 
	 */
	formatResponse := create.FormatTodo(todo)
	c.JSON(http.StatusOK,formatResponse)
}