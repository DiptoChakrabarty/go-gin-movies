package routes

import (
	"github.com/DiptoChakrabarty/go-gin-movies/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieApi struct {
	loginController controller.LoginController
	movieController controller.MovieController
}

func NewMoviesApi(loginController controller.LoginController,
	movieController controller.MovieController) *MovieApi {
	return &MovieApi{
		loginController: loginController,
		movieController: movieController,
	}
}

// All Routes

// Login Route
func (route *MovieApi) LoginMethod(ctx *gin.Context) {
	token := route.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}
}

// Get All Movies Route
func (route *MovieApi) GetAll(ctx *gin.Context) {
	ctx.JSON(200, route.movieController.GetAll())
}

// Add Movie Route
func (route *MovieApi) Add(ctx *gin.Context) {
	err := route.movieController.Add(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Movie Details Saved"})
	}

}

// Update Movie Route
func (route *MovieApi) Update(ctx *gin.Context) {
	err := route.movieController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Movie Details Saved"})
	}

}

// Delete Movie Route
func (route *MovieApi) Delete(ctx *gin.Context) {
	err := route.movieController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Movie Details Saved"})
	}
}

// Check Health of Server
func (route *MovieApi) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Server Up and Running",
	})
}
