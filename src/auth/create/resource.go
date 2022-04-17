package create

import "todosAPI/database"

type ResourceRegistrationUser struct {
	Username string `json:"username"`
}

func FormatRegistration(user database.User) ResourceRegistrationUser {
	formatter := ResourceRegistrationUser{
		Username: user.Username,
	}
	
	return formatter
}