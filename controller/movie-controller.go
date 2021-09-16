package controller

import (
	"github.com/DiptoChakrabarty/go-gin-movies/operation"
	"github.com/gin-gonic/gin"
)

type MovieController interface {
	GetAll() []operation.Movie
	Save(c *gin.Context)
}
