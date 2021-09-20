package service

import (
	"github.com/DiptoChakrabarty/go-gin-movies/models"
	"github.com/DiptoChakrabarty/go-gin-movies/operation"
)

type MovieService interface {
	Add(operation.Movie) error
	Update(operation.Movie) error
	Delete(operation.Movie) error
	GetAll() []operation.Movie
}

type movieService struct {
	model models.MovieModel
}

func New(movieModel models.MovieModel) MovieService {
	return &movieService{
		model: movieModel,
	}
}

func (svc *movieService) Add(movie operation.Movie) error {
	svc.model.Add(movie)
	return nil
}

func (svc *movieService) Update(movie operation.Movie) error {
	svc.model.Update(movie)
	return nil
}

func (svc *movieService) Delete(movie operation.Movie) error {
	svc.model.Delete(movie)
	return nil
}

func (svc *movieService) GetAll() []operation.Movie {
	return svc.model.GetAll()
}
