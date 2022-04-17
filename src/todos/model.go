package todos

import (
	"todosAPI/database"
	"todosAPI/src/todos/update"

	"todosAPI/utils"

	"gorm.io/gorm"
)

type Model interface {
	Read() ([]database.Todos, error)
	Create(todo database.Todos)	(database.Todos, error)
	Update(input update.UpdateTodo,ID string) (database.Todos, error)
	Delete(ID string) (database.Todos, error)
}

type constructor struct {
	db *gorm.DB
}

func InitConstructor() *constructor {
	return &constructor{utils.DB}
}

func (c *constructor) Create(todo database.Todos) (database.Todos, error) {
	// var ctx *gin.Context
	// authPayload := ctx.MustGet("authorization_payload_key").(*token.Payload)
	// todo.Username = authPayload.Username

	err := c.db.Create(&todo).Error
	
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (c *constructor) Read() ([]database.Todos,error) {
	var todo []database.Todos	
	err := c.db.Preload("Category").Order("id").Find(&todo).Error

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (c *constructor) Update(input update.UpdateTodo,ID string) (database.Todos, error) {
	var todo database.Todos
	
	err := c.db.First(&todo,ID).Error
	if err != nil {
		return todo, err
	}

	todo.Title = input.Title
	todo.Description = input.Description
	todo.CategoryID = input.CategoryID
	errSave := c.db.Save(&todo).Error

	if errSave != nil {
		return todo, errSave
	}
	
	return todo, nil
}

func (c *constructor) Delete(ID string) (database.Todos, error) {
	var todo database.Todos
	err := c.db.First(&todo,ID).Delete(&todo).Error
	
	if err != nil {
		return todo, err
	}
	
	return todo, nil
}