package category

import (
	"todosAPI/database"
	"todosAPI/src/category/create"
)

// * membuat interface controller untuk menghandel method2 dengan interface tersebut
type Controller interface {
	GetCategoryController(IsTodo string) ([]database.Category, error)
	FindCategoryController(ID string) (database.Category, error)
	DeleteCategoryController(ID string) (database.Category, error)
	CreateCategoryController(input create.CreateCategoryInput) (database.Category, error)
}

// * memasukan data model yang terisi model
type controller struct {
	model Model
}

// * menginisialisasikan controller dan memasukan ke dalam struct diatas
func InitConstructorController(model Model) *controller {
	return &controller{model}
}

// * membuat function yang sesuai dengan interface diatas
func (c *controller) GetCategoryController(IsTodo string) ([]database.Category, error) {
	// * mengambil controller lalu masuk ke model -> lalu masuk ke method interface model
	categories, err := c.model.Get(IsTodo)
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (c *controller) FindCategoryController(ID string) (database.Category, error) {
	category, err := c.model.Find(ID)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *controller) DeleteCategoryController(ID string) (database.Category, error) {
	category, err := c.model.Delete(ID)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *controller) CreateCategoryController(input create.CreateCategoryInput) (database.Category, error) {
	category := database.Category{}
	category.Name = input.Name
	category.UserID = input.UserID

	newCategory, err := c.model.Create(category)

	if err != nil {
		return category, err
	}

	return newCategory, nil
}
