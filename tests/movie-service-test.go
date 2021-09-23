package tests

import (
	"github.com/DiptoChakrabarty/go-gin-movies/models"
	"github.com/DiptoChakrabarty/go-gin-movies/operation"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

const (
	TITLE       = "The Matrix"
	DESCRIPTION = " Story of the matrix"
	TRAILER     = "https://www.youtube.com/watch?v=AB59tJPthZg"
	PRICE       = 300
	FIRSTNAME   = "Keanu"
	LASTNAME    = "Reeves"
	EMAIL       = "reeves@gmail.com"
	AGE         = 50
)

var testMovie operation.Movie = operation.Movie{
	Title:       TITLE,
	Description: DESCRIPTION,
	Trailer:     TRAILER,
	Price:       PRICE,
	LeadActor: operation.Actor{
		FirstName: FIRSTNAME,
		LastName:  LASTNAME,
		Age:       AGE,
		Email:     EMAIL,
	},
}

var _ = ginkgo.Describe("Movie Service", func() {
	var (
		movieModel   models.MovieModel
		movieService service.MovieService
	)

	ginkgo.BeforeSuite(func() {
		movieModel = models.NewModelDB()
		movieService = service.New(movieModel)
	})

	ginkgo.Describe("Get all movie details", func() {
		ginkgo.Context("If movies in db", func() {
			ginkgo.BeforeEach(func() {
				movieService.Add(testMovie)
			})
			ginkgo.It("atleast one element should return ", func() {
				movies := movieService.GetAll()
				Ω(movies).ShouldNot(gomega.BeEmpty())
			})

			ginkgo.It("display correct fields", func() {
				startMovie := movieService.GetOne(1)
				(startMovie.Title).Should(gomega.Equal(TITLE))
				Ω(startMovie.Description).Should(gomega.Equal(DESCRIPTION))
				Ω(startMovie.Trailer).Should(gomega.Equal(TRAILER))
				Ω(startMovie.Price).Should(gomega.Equal(PRICE))
				Ω(startMovie.LeadActor.Age).Should(gomega.Equal(AGE))
				Ω(startMovie.LeadActor.FirstName).Should(gomega.Equal(FIRSTNAME))
				Ω(startMovie.LeadActor.LastName).Should(gomega.Equal(LASTNAME))
				Ω(startMovie.LeadActor.Email).Should(gomega.Equal(EMAIL))
			})
			ginkgo.AfterEach(func() {
				movie := movieService.GetOne(1)
				movieService.Delete(movie)
			})
		})
		ginkgo.Context("No movie in database", func() {
			ginkgo.It("empty list", func() {
				movies := movieService.GetAll()
				Ω(movies).Should(gomega.BeEmpty())
			})
		})
	})
})
