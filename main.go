package main

import (
	"github.com/DiptoChakrabarty/go-gin-movies/controller"
	"github.com/DiptoChakrabarty/go-gin-movies/middlewares"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"net/http"
)

var (
	movieService    service.MovieService       = service.New()
	movieController controller.MovieController = controller.New(movieService)
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.AuthMethod(), gindump.Dump())

	router.Static("/css", "./templates/css")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server Up and Running",
		})
	})

	movieRoutes := router.Group("/api")
	{
		movieRoutes.GET("/movies", func(ctx *gin.Context) {
			ctx.JSON(200, movieController.GetAll())
		})
		movieRoutes.POST("/movies", func(ctx *gin.Context) {
			err := movieController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Movie Details Saved"})
			}

		})
	}

	viewRoutes := router.Group("/view")
	{
		viewRoutes.GET("/movies", func(ctx *gin.Context) {
			movieController.DisplayAll(ctx)
		})
	}

	router.Run(":8000")
}
