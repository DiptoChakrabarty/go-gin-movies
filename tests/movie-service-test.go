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
				gomega.Expect(movies).ToNot(gomega.BeEmpty())
			})

			ginkgo.It("display correct fields", func() {
				startMovie := movieService.GetAll()[0]
				gomega.Expect(startMovie.Title).To(gomega.Equal(TITLE))
				gomega.Expect(startMovie.Description).To(gomega.Equal(DESCRIPTION))
				gomega.Expect(startMovie.Trailer).To(gomega.Equal(TRAILER))
				gomega.Expect(startMovie.Price).To(gomega.Equal(PRICE))
				gomega.Expect(startMovie.LeadActor.Age).To(gomega.Equal(AGE))
				gomega.Expect(startMovie.LeadActor.FirstName).To(gomega.Equal(FIRSTNAME))
				gomega.Expect(startMovie.LeadActor.LastName).To(gomega.Equal(LASTNAME))
				gomega.Expect(startMovie.LeadActor.Email).To(gomega.Equal(EMAIL))
			})
			ginkgo.AfterEach(func() {
				movie := movieService.GetOne(1)
				movieService.Delete(movie)
			})

		})
		ginkgo.Context("No movie in database", func() {
			ginkgo.It("empty list", func() {
				movies := movieService.GetAll()
				gomega.Expect(movies).To(gomega.BeEmpty())
			})
		})
	})
})
