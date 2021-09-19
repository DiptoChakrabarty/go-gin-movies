package middlewares

import (
	"log"
	"net/http"

	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthoirzeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenUser := authHeader[len(SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenUser)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Name : ", claims["name"])
			log.Println("Issuer : ", claims["issuer"])
			log.Println("Admin : ", claims["admin"])
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
