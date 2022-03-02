package v1

import (
	"lottery-web-scrapping/controllers"
	"lottery-web-scrapping/middlewares"

	"github.com/gin-gonic/gin"
)

func LotteryRouter(r *gin.RouterGroup) {
	r.Use(middlewares.AuthorizationAPIKey())
	{
		// Check Drawing By Date
		r.GET("/:DateId", controllers.FindLotteryByDate)

		// Check Drawing By Date and Number
		r.GET("/:DateId/:NumberId", controllers.FindLotteryByNumber)

		// Check Drawing By Date and Numbers
		r.POST("/:DateId")

	}
}
