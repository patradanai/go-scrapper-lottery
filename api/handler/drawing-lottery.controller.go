package handler

import (
	httpError "lottery-web-scrapping/pkg/http-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paramsDate struct {
	DateID string `uri:"DateId" binding:"required,len=8"`
}

type paramNumber struct {
	DateID string `uri:"DateId" binding:"required,len=8"`
	Number string `uri:"NumberId" binding:"required,alphanum,len=6"`
}

func (h *Handler) FindLotteryByDate(c *gin.Context) {

	// Binding
	params := paramsDate{}
	if err := c.ShouldBindUri(&params); err != nil {
		httpError.NewBadRequest(c, "please Check parameter uri")
		return
	}

	if err := h.drawingDateService.FindDrawingDate(params.DateID); err != nil {
		httpError.NewBadRequest(c, "drawing Date not existing")
		return
	}
	// Find Drawing By Date
	result, err := h.lotteryService.FindLotteryByDate(params.DateID)
	if err != nil {
		httpError.NewInternalServerError(c, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "drawing_date": params.DateID, "result": result})
}

func (h *Handler) FindLotteryByNumber(c *gin.Context) {

	// Binding
	params := paramNumber{}
	if err := c.ShouldBindUri(&params); err != nil {
		httpError.NewBadRequest(c, "please Check parameter uri")
		return
	}

	if err := h.drawingDateService.FindDrawingDate(params.DateID); err != nil {
		httpError.NewBadRequest(c, "drawing Date not existing")
		return
	}

	// Find Number by Date
	result, exist := h.lotteryService.FindLotteryByNumber(params.Number, params.DateID)
	if exist {
		c.JSON(http.StatusOK, gin.H{"success": true, "drawing_date": params.DateID, "winner": true, "lottery": params.Number, "result": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "drawing_date": params.DateID, "winner": false, "lottery": params.Number, "message": "โดนหวยแดก"})
}

func (h *Handler) FindLotteriesDate(c *gin.Context) {

}
