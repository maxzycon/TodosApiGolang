package create

import "todosAPI/database"

type ResourceRegistrationUser struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

func FormatRegistration(user database.User) ResourceRegistrationUser {
	formatter := ResourceRegistrationUser{
		ID: int(user.ID),
		Username: user.Username,
	}
	
	return formatter
}