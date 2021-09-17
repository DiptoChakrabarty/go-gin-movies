package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMethod() gin.HandleFunc {
	return gin.BasicAuth(gin.Accounts{
		"pinku": "qwerty"
	})
}