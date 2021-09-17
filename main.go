package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DiptoChakrabarty/go-gin-movies/middlewares"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/DiptoChakrabarty/go-gin-movies/controller"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	movieService service.MovieService = service.New()
	movieController controller.MovieController = controller.New(movieService)
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery(),middlewares.Logger(),
	middlewares.AuthMethod,gindump.Dump())

	router.GET("/health",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message": "Server Up and Running"
		})
	})

	router.GET("/movies",func(ctx *gin.Context){
		ctx.JSON(200,movieController.FindAll())
	})
	router.POST("/movies",func(ctx *gin.Context){
		err := movieController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK,gin.H{"message": "Movie Details Saved"})
		}
		
	})

	router.Run(":8000")
}