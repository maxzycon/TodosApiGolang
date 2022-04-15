package routes

import (
	"net/http"
	"todosAPI/todos"
	"todosAPI/todos/create"
	"todosAPI/todos/read"

	"github.com/gin-gonic/gin"
)

type todoRoute struct {
	todoCreate todos.Controller
}

func TodoRouteInit(todoController todos.Controller) *todoRoute  {
	return &todoRoute{todoController}
}

func (r *todoRoute) CreateTodo(c *gin.Context) {
	var input create.CreateTodoInput
	err := c.ShouldBindJSON(&input)
	
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	todos,err_response := r.todoCreate.CreateTodo(input)
	
	if err_response != nil {
		c.JSON(http.StatusUnprocessableEntity,err_response)
		return
	}
	formatResponse := create.FormatTodo(todos)
	c.JSON(http.StatusOK,formatResponse)
}

func (r *todoRoute) ReadTodo(c *gin.Context) {
	todos,err := r.todoCreate.ReadTodo()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity,err)
		return
	}
	formatResponse := read.FormatTodos(todos)
	c.JSON(http.StatusOK,formatResponse)
}