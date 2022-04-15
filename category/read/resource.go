package read

import "todosAPI/database"

type FormatterCategoryCreate struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Todos []TodoFormatterCreate `json:"todos"`
}

type TodoFormatterCreate struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"desc"`
}

func FormatCategory(category database.Category) FormatterCategoryCreate {
	formatter := FormatterCategoryCreate{
		ID: category.ID,
		Name: category.Name,
	}
	return formatter
}

func FormatCategories(categories []database.Category) []FormatterCategoryCreate {
	var categoriesformatter []FormatterCategoryCreate

	for _, category := range categories{
		// init todosformatter 
		todosformatter := []TodoFormatterCreate{}
		categoryformatter := FormatCategory(category)
		categoryformatter.Todos = todosformatter

		if len(category.Todos) != 0 {
			// foreach semua todos
			for _, todo := range category.Todos {
				// init todoFormatter
				todosData := TodoFormatterCreate{}
				// masukan todo ke todosFormatter
				todosData.ID = todo.ID
				todosData.Title = todo.Title
				todosData.Description = todo.Description
				todosformatter = append(todosformatter,todosData)
			}
			categoryformatter.Todos = todosformatter
		}

		categoriesformatter = append(categoriesformatter,categoryformatter)
	}
	return categoriesformatter
}