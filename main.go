package main

import (
	"github.com/gin-gonic/gin"
)

var (
	movieService service.MovieService = service.New()
	movieController controller.MovieController = controller.New(movieService)
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery(),middlewares.Logger())

	router.GET("/health",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message": "Server Up and Running"
		})
	})

	router.GET("/movies",func(ctx *gin.Context){
		ctx.JSON(200,movieController.FindAll())
	})
	router.POST("/movies",func(ctx *gin.Context){
		ctx.JSON(200,movieController.Save(ctx))
	})

	router.Run(":8000")
}