package service

import (
	"github.com/DiptoChakrabarty/go-gin-movies/models"
	"github.com/DiptoChakrabarty/go-gin-movies/user"
)

type LoginService interface {
	Login(username string, password string) bool
	Register(newuser user.User) error
	GetAllUsers() []user.User
}

type loginService struct {
	model models.MovieModel
}

func NewLoginService(movieModel models.MovieModel) LoginService {
	return &loginService{
		model: movieModel,
	}
}

func (svc *loginService) Login(username string, password string) bool {
	verify_user := svc.model.GetUser(username)
	if (user.User{}) == verify_user {
		return false
	}
	return verify_user.PassWord == password
}

func (svc *loginService) Register(newuser user.User) error {
	svc.model.AddUser(newuser)
	return nil
}

func (svc *loginService) GetAllUsers() []user.User {
	return svc.model.GetAllUsers()
}
