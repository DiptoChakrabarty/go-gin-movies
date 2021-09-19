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
	movieService service.MovieService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	movieController controller.MovieController  = controller.New(movieService)
	loginController controlller.LoginController = controlller.NewLoginController(loginService, jwtService)
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

	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}

	})

	movieRoutes := router.Group("/api", middlewares.AuthoirzeUser())
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
