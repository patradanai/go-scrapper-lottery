package v1

import (
	"lottery-web-scrapping/api/handler"
	"lottery-web-scrapping/api/middlewares"

	"github.com/gin-gonic/gin"
)

func OauthRouter(r *gin.RouterGroup, handler *handler.Handler) {
	r.GET("/oauth", middlewares.AuthorizationJWT())
	r.POST("/oauth", middlewares.AuthorizationJWT(), handler.CreateOAuth)
	r.PATCH("/oauth/:oauthId", middlewares.AuthorizationJWT())
}
