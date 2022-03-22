package v1

import (
	"lottery-web-scrapping/api/handler"
	"lottery-web-scrapping/api/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup, handler *handler.Handler) {
	r.POST("/signup", handler.SignUp)
	r.POST("/signin", middlewares.Authentication())
	r.POST("/signout")
	r.POST("/refresh", handler.RefreshToken)
}
