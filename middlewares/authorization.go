package middlewares

import "github.com/gin-gonic/gin"

func AuthorizationJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func AuthorizationAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
