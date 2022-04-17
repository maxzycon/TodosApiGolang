package database

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string
	UserID string `gorm:"type:varchar;size:255"`
	User User `gorm:"references:username"`
	Todos []Todos
}

