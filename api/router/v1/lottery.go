package v1

import (
	"lottery-web-scrapping/api/handler"
	"lottery-web-scrapping/api/middlewares"

	"github.com/gin-gonic/gin"
)

func LotteryRouter(r *gin.RouterGroup, handler *handler.Handler) {
	r.Use(middlewares.AuthorizationAPIKey())
	{
		// Check Drawing By Date
		r.GET("/:DateId", handler.FindLotteryByDate)

		// Check Drawing By Date and Number
		r.GET("/:DateId/:NumberId", handler.FindLotteryByNumber)

		// Check Drawing By Date and Numbers
		r.POST("/:DateId")

	}
}
