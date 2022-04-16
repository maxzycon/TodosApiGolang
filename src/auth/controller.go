package auth

import (
	"todosAPI/database"
	"todosAPI/src/auth/create"
)

type Controller interface {
	CreateUserController(input create.RegistrationUserInput) (database.User, error)
}

type controller struct {
	model Model
}

func InitConstructorController(model Model) *controller {
	return &controller{model}
}

func (c *controller) CreateUserController(input create.RegistrationUserInput) (database.User, error) {
	var user database.User
	
	user.Username = input.Username
	// TODO: create hash_password
	user.Hash_password = input.Password
	user,err := c.model.Create(user)

	if err != nil {
		return user, err
	}

	return user, nil
}
