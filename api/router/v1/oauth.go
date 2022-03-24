package v1

import (
	"lottery-web-scrapping/api/handler"
	"lottery-web-scrapping/api/middlewares"

	"github.com/gin-gonic/gin"
)

func OauthRouter(r *gin.RouterGroup, handler *handler.Handler) {
	r.Use(middlewares.AuthorizationJWT())
	{
		r.GET("/oauth", handler.GetOAuth)
		r.POST("/oauth", handler.CreateOAuth)
		r.PATCH("/oauth/:oauthId", handler.UpdateOAuth)
	}
}
