package auth

import (
	"todosAPI/database"
	"todosAPI/utils"

	"gorm.io/gorm"
)

type Model interface {
	Create(user database.User) (database.User, error)
	FindUsername(user database.User) (database.User, error)
}

type constructor struct {
	db *gorm.DB
}

func InitConstructor() *constructor {
	return &constructor{utils.DB}
}

func (m *constructor) Create(user database.User) (database.User, error) {
	err := m.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *constructor) FindUsername(user database.User) (database.User, error) {
	err := m.db.Where("username = ?",user.Username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
