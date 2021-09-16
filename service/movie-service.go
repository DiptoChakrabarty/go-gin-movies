package service

import (
	"github.com/DiptoChakrabarty/go-gin-movies/operation"
)

type MovieService interface {
	Save(operation.Movie) operation.Movie
	GetAll() []operation.Movie
}

type movieService struct {
	movies []operation.Movie
}

func New() MovieService {
	return &movieService{}
}

func (svc *movieService) Save(movie operation.Movie) []operation.Movie {
	svc.movies = append(svc.movies, movie)
	return movie
}

func (svc *movieService) GetAll() []operation.Movie {
	return svc.movies
}
