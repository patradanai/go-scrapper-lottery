package routes

import (
	v1 "lottery-web-scrapping/routes/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitialRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "ROOT WEB")
		return
	})

	r.POST("/signin")
	r.POST("/signout")
	r.POST("/refresh")

	// Api V1
	v1API := r.Group("/api/v1")
	{
		// Auth
		v1.AuthRouter(v1API)
	}

	return r
}
