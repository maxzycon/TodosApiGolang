package category

import (
	"todosAPI/database"
	"todosAPI/src/category/create"
)
type Controller interface {
	GetCategoryController() ([]database.Category, error)
	FindCategoryController(ID string) (database.Category, error)
	DeleteCategoryController(ID string) (database.Category, error)
	CreateCategoryController(input create.CreateCategoryInput) (database.Category, error)
}

type controller struct {
	model Model
}

func InitConstructorController(model Model) *controller {
	return &controller{model}
}

func (c *controller) GetCategoryController() ([]database.Category, error)  {
	categories,err := c.model.Get()
	if err != nil {
		return categories,err
	}
	return categories,nil
}

func (c *controller) FindCategoryController(ID string) (database.Category, error)  {
	category,err := c.model.Find(ID)
	if err != nil {
		return category,err
	}
	return category,nil
}

func (c *controller) DeleteCategoryController(ID string) (database.Category, error)  {
	category,err := c.model.Delete(ID)
	if err != nil {
		return category,err
	}
	return category,nil
}

func (c *controller) CreateCategoryController(input create.CreateCategoryInput) (database.Category, error)  {
	category := database.Category{}
	category.Name = input.Name

	newCategory,err := c.model.Create(category)
	if err != nil {
		return category,err
	}

	return newCategory,nil
}