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

// Authenticate godoc
// @Summary Generates a JWT
// @ID Authentication
// @Consume application/json
// @Produce json
// @Param user body user.User true "Login User"
// @Success 200
// @Failure 401
// @Router /auth/login [post]
// @Description Authenticates a user and provides a JWT to Authorize API calls
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

// Register godoc
// @Summary Registers a User
// @Consume application/json
// @Produce json
// @Param user body user.User true "Register User"
// @Success 200
// @Failure 401
// @Router /auth/register [post]
// @Description Registers a User in the application
func (route *MovieApi) RegisterMethod(ctx *gin.Context) {
	err := route.loginController.Register(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User Registered Succesfully"})
	}
}

// Users godoc
// @Summary Displays all Users
// @Produce json
// @Success 200
// @Failure 401
// @Router /auth/userdisplay [get]
// @Description Displays all users
func (route *MovieApi) GetAllUsersMethod(ctx *gin.Context) {
	ctx.JSON(200, route.loginController.GetAllUsers(ctx))
}

// GetMovies godoc
// @Security bearerAuth
// @Summary Displays all Movies Info
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /movies [get]
// @Description Gets all the existing Movies from database
func (route *MovieApi) GetAll(ctx *gin.Context) {
	ctx.JSON(200, route.movieController.GetAll())
}

// AddMovies godoc
// @Security bearerAuth
// @Summary Adds a Movie Info
// @Accept json
// @Produce json
// @Param movie body operation.Movie true "Add Movie"
// @Success 200
// @Failure 401
// @Failure 400
// @Router /movies [post]
// @Description Adds a movie to database
func (route *MovieApi) Add(ctx *gin.Context) {
	err := route.movieController.Add(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Movie Details Saved"})
	}

}

// UpdateMovies godoc
// @Security bearerAuth
// @Summary Updates a Movie Info
// @Accept json
// @Produce json
// @Param id path int true "Update Movie"
// @Success 200
// @Failure 401
// @Failure 400
// @Router /movies/{id} [put]
// @Description Updates a movie in database
func (route *MovieApi) Update(ctx *gin.Context) {
	err := route.movieController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Movie Details Saved"})
	}

}

// DeleteMovies godoc
// @Security bearerAuth
// @Summary Deletes a Movie Info
// @Produce json
// @Param id path int true "Delete Movie"
// @Success 200
// @Failure 401
// @Failure 400
// @Router /movies/{id} [delete]
// @Description Deletes a movie in database
func (route *MovieApi) Delete(ctx *gin.Context) {
	err := route.movieController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Movie Details Saved"})
	}
}

// GetMovie godoc
// @Security bearerAuth
// @Summary Displays one movie detail
// @Produce json
// @Param id path int true "Display Movie"
// @Success 200
// @Failure 401
// @Failure 400
// @Router /movies/{id} [get]
// @Description Displays a movie in database
func (route *MovieApi) GetOne(ctx *gin.Context) {
	ctx.JSON(200, route.movieController.GetOne(ctx))
}

// Check Health of Server
// HealthCheck godoc
// @Summary Checks Health of Server
// @Produce json
// @Success 200
// @Failure 401
// @Failure 400
// @Router /health [get]
// @Description Checks Server Health While Running
func (route *MovieApi) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Server Up and Running",
	})
}
