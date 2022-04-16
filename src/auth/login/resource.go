package login

import "todosAPI/database"

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

func NewUser(user database.User) User {
	formatter := User{
		ID: int(user.ID),
		Username: user.Username,
	}
	return formatter
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	User User `json:"user"`
}