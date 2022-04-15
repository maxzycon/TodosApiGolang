package create

import "todosAPI/database"

type FormatterCategoryCreate struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

func FormatCategory(category database.Category) FormatterCategoryCreate {
	formatter := FormatterCategoryCreate{
		ID: category.ID,
		Name: category.Name,
	}
	return formatter
}