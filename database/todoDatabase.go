package database

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	Title string
	Description string
	CategoryID int
	UserID string `gorm:"type:varchar;size:255"`
	User User `gorm:"references:username"`
	Category Category
}

