package routes

import (
	"net/http"
	"todosAPI/src/auth"
	"todosAPI/src/auth/create"

	"github.com/gin-gonic/gin"
)

type authRoute struct {
	authController auth.Controller
}

func AuthRouteInit(authController auth.Controller) *authRoute  {
	return &authRoute{authController}
}

func (r *authRoute) RegistrationUserRoute(c *gin.Context) {
	var input create.RegistrationUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	user,err_response := r.authController.CreateUserController(input)

	if err_response != nil {
		c.JSON(http.StatusUnprocessableEntity,err_response)
		return
	}
	formatter := create.FormatRegistration(user)
	c.JSON(http.StatusOK,formatter)
}