package auth

import (
	"todosAPI/database"
	"todosAPI/src/auth/create"
	"todosAPI/src/auth/login"
	"todosAPI/utils"
)

type Controller interface {
	CreateUserController(input create.RegistrationUserInput) (database.User, error)
	LoginController(input login.LoginUserInput) (database.User, error)
}

type controller struct {
	model Model
}

func InitConstructorController(model Model) *controller {
	return &controller{model}
}

func (c *controller) CreateUserController(input create.RegistrationUserInput) (database.User, error) {
	user  := database.User{}
	
	user.Username = input.Username
	hashpassword, herr := utils.HashPassword(input.Password)
	
	if herr != nil {
		return user,herr
	}
	
	user.Hash_password = hashpassword
	user,err := c.model.Create(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (c *controller) LoginController(input login.LoginUserInput) (database.User, error) {
	user  := database.User{}
	user.Username = input.Username
	user.Hash_password = input.Password
	user, err := c.model.FindUsername(user)

	if err != nil {
		return user, err
	}

	errCheck := utils.CheckPassword(input.Password,user.Hash_password)

	if errCheck != nil {
		return user, errCheck
	}
	
	return user, nil
}
