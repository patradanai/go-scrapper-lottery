package v1

import (
	"lottery-web-scrapping/api/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/singup")
	r.POST("/signin", middlewares.Authentication())
	r.POST("/signout")
	r.POST("/refresh")
}
