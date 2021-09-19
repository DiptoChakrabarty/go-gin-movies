package controlller

import (
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/DiptoChakrabarty/go-gin-movies/user"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (l *loginController) Login(ctx *gin.Context) string {
	var UserCredentials user.User
	err := ctx.ShouldBind(&UserCredentials)
	if err != nil {
		return ""
	}
	isAuth := l.loginService.Login(UserCredentials.UserName, UserCredentials.PassWord)
	if isAuth {
		return l.jwtService.GenerateToken(UserCredentials.UserName, true)
	}
	return ""
}
