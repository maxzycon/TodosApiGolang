package create

import (
	"todosAPI/database"
)

type FormatterTodosCreate struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"desc"`
	Category FormatterCategory `json:"category"`
}

type FormatterCategory struct {
	ID uint `json:"id"`
	Name  string `json:"name"`
}

func FormatTodo(todo database.Todos) FormatterTodosCreate {
	formatter := FormatterTodosCreate{
		ID:todo.ID,
		Title: todo.Title,
		Description: todo.Description,
	}
	return formatter
}

func FormatTodos(todos []database.Todos) []FormatterTodosCreate {
	// inisialisasi stuct
	var todosformatter []FormatterTodosCreate
	for _, todo := range todos{
		// format todo 
		todoformatter := FormatTodo(todo)
		// eager loading category
		category := todo.Category
		// initsialisasi struct
		todosCategory := FormatterCategory{}
		// masukan data dari eager loading ke struct
		todosCategory.ID = category.ID
		todosCategory.Name = category.Name
		todoformatter.Category = todosCategory
		// menambahkan object ke array
		todosformatter = append(todosformatter,todoformatter)
	}
	return todosformatter
}
