package controller

import (
	"net/http"

	"github.com/DiptoChakrabarty/go-gin-movies/operation"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/DiptoChakrabarty/go-gin-movies/validators"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type MovieController interface {
	GetAll() []operation.Movie
	Save(ctx *gin.Context) error
	DisplayAll(ctx *gin.Context)
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

func (c *controller) DisplayAll(ctx *gin.Context) {
	movies := c.svc.GetAll()
	data := gin.H{
		"title":  "Movies",
		"movies": movies,
	}

	ctx.HTML(http.StatusOK, "main.html", data)
}
