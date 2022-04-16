package routes

import (
	"net/http"
	"time"
	"todosAPI/src/auth"
	"todosAPI/src/auth/create"
	"todosAPI/src/auth/login"
	"todosAPI/src/token"

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

func (r *authRoute) LoginUserRoute(c *gin.Context) {
	var input login.LoginUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	user,err_response := r.authController.LoginController(input)

	if err_response != nil {
		c.JSON(http.StatusUnprocessableEntity,err_response.Error())
		return
	}
	waktu,_ := time.ParseDuration("15m")

	token,_ := token.NewJWTMaker("12345678901234567890123456789012")
	
	accessToken,err := token.CreateToken(
		user.Username,
		waktu,
	)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity,err.Error())
		return 
	}

	rsp := login.LoginUserResponse{
		AccessToken: accessToken,
		User: login.NewUser(user),
	}

	c.JSON(http.StatusOK,rsp)
}