package todos

import (
	"todosAPI/database"

	"todosAPI/utils"

	"gorm.io/gorm"
)

type Model interface {
	Read() ([]database.Todos, error)
	Create(todo database.Todos)	(database.Todos, error)
}

type constructor struct {
	db *gorm.DB
}

func InitConstructor() *constructor {
	return &constructor{utils.DB}
}

func (c *constructor) Create(todo database.Todos) (database.Todos, error) {
	err := c.db.Create(&todo).Error
	
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (c *constructor) Read() ([]database.Todos,error) {
	var todo []database.Todos
	err := c.db.Preload("Category").Find(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}
