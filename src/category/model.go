package category

import (
	"todosAPI/database"
	"todosAPI/utils"

	"gorm.io/gorm"
)

type Model interface {
	Get(IsTodo string) ([]database.Category,error)
	Find(ID string) (database.Category,error) 
	Delete(ID string) (database.Category,error) 
	Create(category database.Category)	(database.Category, error)
}

type constructor struct {
	db *gorm.DB
}

func InitConstructor() *constructor {
	return &constructor{utils.DB}
}

func (c *constructor) Get(IsTodo string) ([]database.Category,error) {
	var dataCategory []database.Category
	if IsTodo == "1" {
		c.db.Preload("Todos").Find(&dataCategory)
	}else{
		c.db.Find(&dataCategory)
	}
	return dataCategory,nil
}

func (c *constructor) Create(category database.Category) (database.Category, error) {
	err := c.db.Create(&category).Error
	
	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *constructor) Find(ID string) (database.Category, error) {
	var category database.Category
	err := c.db.First(&category,ID).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *constructor) Delete(ID string) (database.Category, error) {
	var category database.Category
	err := c.db.Delete(&category,ID).Error

	if err != nil {
		return category, err
	}

	return category, nil
}