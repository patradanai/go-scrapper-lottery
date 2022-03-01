package v1

import (
	"lottery-web-scrapping/middlewares"

	"github.com/gin-gonic/gin"
)

func LotteryRouter(r *gin.RouterGroup) {
	r.Use(middlewares.AuthorizationAPIKey())
	{
		// Check Drawing By Date
		r.GET("/:Date")

		// Check Drawing By Date and Number
		r.GET("/:Date/:Number")

		// Check Drawing By Date and Numbers
		r.POST("/:Date")

	}
}
