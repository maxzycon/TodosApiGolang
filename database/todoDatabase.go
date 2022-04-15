package database

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	Title string
	Description string
	CategoryID int
	Category Category
}

