package controller

import (
	"github.com/DiptoChakrabarty/go-gin-movies/operation"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/gin-gonic/gin"
)

type MovieController interface {
	GetAll() []operation.Movie
	Save(c *gin.Context) operation.Movie
}

type controller struct {
	svc service.MovieService
}

func New(svc service.MovieService) MovieController {
	return controller{
		svc: svc,
	}
}

func (c *controller) GetAll() []operation.Movie {
	return c.svc.GetAll()
}

func (c *controller) Save(ctx *gin.Context) operation.Movie {
	var movie operation.Movie
	ctx.BindJSON(&movie)
	c.svc.Save(movie)
	return movie
}
