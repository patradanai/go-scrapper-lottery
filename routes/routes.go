package routes

import (
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
	r.POST("/auth")

	// Api V1
	v1 := r.Group("/api/v1")

	v1.Use()
	{

	}

	return r
}
