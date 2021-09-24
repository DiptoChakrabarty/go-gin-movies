package controller

import (
	"strconv"

	"github.com/DiptoChakrabarty/go-gin-movies/operation"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/DiptoChakrabarty/go-gin-movies/validators"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type MovieController interface {
	GetAll() []operation.Movie
	Add(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	GetOne(ctx *gin.Context) operation.Movie
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

func (c *controller) Add(ctx *gin.Context) error {
	var movie operation.Movie
	err := ctx.ShouldBindJSON(&movie)
	if err != nil {
		return err
	}
	err = validate.Struct(movie)
	if err != nil {
		return err
	}
	c.svc.Add(movie)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var movie operation.Movie
	err := ctx.ShouldBindJSON(&movie)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	movie.ID = id

	err = validate.Struct(movie)
	if err != nil {
		return err
	}
	c.svc.Update(movie)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var movie operation.Movie

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	movie.ID = id

	c.svc.Delete(movie)
	return nil
}

func (c *controller) GetOne(ctx *gin.Context) operation.Movie {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	return c.svc.GetOne(id)
}
