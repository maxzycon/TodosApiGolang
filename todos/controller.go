package todos

import (
	"todosAPI/database"
	"todosAPI/todos/create"
	"todosAPI/todos/update"
)
type Controller interface {
	CreateTodo(input create.CreateTodoInput) (database.Todos, error)
	ReadTodo() ([]database.Todos, error)
	UpdateTodo(input update.UpdateTodo,id string) (database.Todos, error)
}

type controller struct {
	model Model
}

func InitConstructorController(model Model) *controller {
	return &controller{model}
}

func (c *controller) CreateTodo(input create.CreateTodoInput) (database.Todos, error)  {
	todo := database.Todos{}
	todo.Title = input.Title
	todo.Description = input.Description
	todo.CategoryID = input.CategoryID

	newTodo,err := c.model.Create(todo)
	if err != nil {
		return todo,err
	}

	return newTodo,nil
}

func (c *controller) ReadTodo() ([]database.Todos, error) {
	todos,err := c.model.Read()
	if err != nil {
		return todos,err
	}

	return todos,nil
}

func (c *controller) UpdateTodo(input update.UpdateTodo,id string) (database.Todos, error) {
	todo,err := c.model.Update(input,id)
	if err != nil {
		return todo,err
	}

	return todo,nil
}