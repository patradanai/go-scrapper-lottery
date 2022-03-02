package controllers

import (
	"fmt"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/repositories"
	"lottery-web-scrapping/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ParamsDate struct {
	DateID string `uri:"DateId" binding:"required,len=8"`
}

type ParamNumber struct {
	DateID string `uri:"DateId" binding:"required,len=8"`
	Number string `uri:"NumberId" binding:"required,alphanum,len=6"`
}

func FindLotteryByDate(c *gin.Context) {
	LotteryRepo := repositories.NewDrawingLotteryRepository(configs.ClientMongo)
	LotteryService := services.NewDrawingLotteryService(LotteryRepo)

	params := ParamsDate{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Please Check parameter uri"})
		return
	}
	fmt.Println(params.DateID)
	result, err := LotteryService.FindLotteryByDate(params.DateID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Somthing went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "drawing_date": params.DateID, "result": result})
}

func FindLotteryByNumber(c *gin.Context) {
	LotteryRepo := repositories.NewDrawingLotteryRepository(configs.ClientMongo)
	LotteryService := services.NewDrawingLotteryService(LotteryRepo)

	params := ParamNumber{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Please Check parameter uri"})
		return
	}
	fmt.Printf("%v,%v", params.DateID, params.Number)
	result, exist := LotteryService.FindLotteryByNumber(params.Number, params.DateID)
	if exist {
		c.JSON(http.StatusAccepted, gin.H{"success": true, "drawing_date": params.DateID, "winner": true, "lottery": params.Number, "result": result})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "drawing_date": params.DateID, "winner": false, "lottery": params.Number, "message": "โดนหวยแดก"})
}
