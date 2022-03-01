package v1

import "github.com/gin-gonic/gin"

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/singup")
	r.POST("/signin")
	r.POST("/signout")
	r.POST("/refresh")
}
