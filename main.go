package main

import (
	"github.com/DiptoChakrabarty/go-gin-movies/controller"
	"github.com/DiptoChakrabarty/go-gin-movies/docs"
	"github.com/DiptoChakrabarty/go-gin-movies/middlewares"
	"github.com/DiptoChakrabarty/go-gin-movies/models"
	"github.com/DiptoChakrabarty/go-gin-movies/routes"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoModel   models.MovieModel    = models.NewModelDB()
	movieService service.MovieService = service.New(videoModel)
	loginService service.LoginService = service.NewLoginService(videoModel)
	jwtService   service.JWTService   = service.NewJWTService()

	movieController controller.MovieController = controller.New(movieService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization

func main() {

	docs.SwaggerInfo.Title = "MovieS API"
	docs.SwaggerInfo.Description = "This is a Movies Api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.New()
	router.Use(gin.Recovery())

	movieAPI := routes.NewMoviesApi(loginController, movieController)

	movieRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		login := movieRoutes.Group("/auth")
		{
			login.POST("/register", movieAPI.RegisterMethod)
			login.POST("/login", movieAPI.LoginMethod)
			login.GET("/userdisplay", movieAPI.GetAllUsersMethod)
		}

		movies := movieRoutes.Group("/movies", middlewares.AuthoirzeUser())
		{
			movies.GET("", movieAPI.GetAll)
			movies.POST("", movieAPI.Add)
			movies.PUT(":id", movieAPI.Update)
			movies.DELETE(":id", movieAPI.Delete)
			movies.GET(":id", movieAPI.GetOne)
		}

		health := movieRoutes.Group("/health")
		{
			health.GET("", movieAPI.HealthCheck)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8000")
}
