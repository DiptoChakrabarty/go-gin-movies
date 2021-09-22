package middlewares

import (
	//"fmt"
	"github.com/DiptoChakrabarty/go-gin-movies/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthoirzeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const SCHEMA = "Bearer "
		//fmt.Println(SCHEMA)
		authHeader := ctx.GetHeader("Authorization")
		//fmt.Println(authHeader)
		tokenUser := authHeader[len(SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenUser)

		if token.Valid {
			_ = token.Claims.(jwt.MapClaims)
			//log.Println("Name : ", claims["name"])
			//log.Println("Issuer : ", claims["issuer"])
			//log.Println("Admin : ", claims["admin"])
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
