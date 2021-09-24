package controller

import (
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/DiptoChakrabarty/go-gin-movies/user"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
	Register(ctx *gin.Context) error
	GetAllUsers(ctx *gin.Context) []user.User
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

func (l *loginController) Register(ctx *gin.Context) error {
	var NewUser user.User
	err := ctx.ShouldBind(&NewUser)
	if err != nil {
		panic(err)
	}
	err = l.loginService.Register(NewUser)
	return err
}

func (l *loginController) GetAllUsers(ctx *gin.Context) []user.User {
	return l.loginService.GetAllUsers()
}
