package router

import (
	"lottery-web-scrapping/api/handler"
	v1 "lottery-web-scrapping/api/router/v1"

	"github.com/gin-gonic/gin"
)

func InitialRouter(handler *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Api V1
	v1API := r.Group("/api/v1")
	{
		// Auth
		v1.AuthRouter(v1API)

		// Lotery
		v1.LotteryRouter(v1API.Group("/lottery"), handler)
	}

	return r
}
