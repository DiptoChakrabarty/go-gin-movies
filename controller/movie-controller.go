package controller

import (
	"github.com/DiptoChakrabarty/go-gin-movies/operation"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/DiptoChakrabarty/go-gin-movies/validators"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type MovieController interface {
	GetAll() []operation.Movie
	Save(c *gin.Context) error
}

type controller struct {
	svc service.MovieService
}

func New(svc service.MovieService) MovieController {
	validate = validator.New()
	validate.RegisterValidation("harrypotter-check", validators.ValidateMovie)
	return &controller{
		svc: svc,
	}
}

func (c *controller) GetAll() []operation.Movie {
	return c.svc.GetAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var movie operation.Movie
	err := ctx.ShouldBindJSON(&movie)
	if err != nil {
		return err
	}
	err = validate.Struct(movie)
	if err != nil {
		return err
	}
	c.svc.Save(movie)
	return nil
}
