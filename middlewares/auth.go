package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMethod() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"pinku": "qwerty",
	})
}
