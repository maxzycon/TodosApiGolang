package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Hash_password string
}