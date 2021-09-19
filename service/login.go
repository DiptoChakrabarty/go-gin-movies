package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedName     string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedName:     "pinku",
		authorizedPassword: "qwerty",
	}
}

func (svc *loginService) Login(username string, password string) bool {
	return svc.authorizedName == username && svc.authorizedPassword == password
}
