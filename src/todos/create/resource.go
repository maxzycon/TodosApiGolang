package create

import (
	"todosAPI/database"
)

type FormatterTodosCreate struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"desc"`
}

func FormatTodo(todo database.Todos) FormatterTodosCreate {
	formatter := FormatterTodosCreate{
		ID:todo.ID,
		Title: todo.Title,
		Description: todo.Description,
	}
	return formatter
}