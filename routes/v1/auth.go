package v1

import (
	"github.com/gin-gonic/gin"
	"lottery-web-scrapping/middlewares"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/singup")
	r.POST("/signin", middlewares.Authentication())
	r.POST("/signout")
	r.POST("/refresh")
}
