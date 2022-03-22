package handler

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Please Check parameter uri"})
		return
	}

	if err := h.drawingDateService.FindDrawingDate(params.DateID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Drawing Date not existing"})
		return
	}
	// Find Drawing By Date
	result, err := h.lotteryService.FindLotteryByDate(params.DateID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Somthing went wrong"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "drawing_date": params.DateID, "result": result})
}

func (h *Handler) FindLotteryByNumber(c *gin.Context) {

	// Binding
	params := paramNumber{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Please Check parameter uri"})
		return
	}

	if err := h.drawingDateService.FindDrawingDate(params.DateID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Drawing Date not existing"})
		return
	}

	// Find Number by Date
	result, exist := h.lotteryService.FindLotteryByNumber(params.Number, params.DateID)
	if exist {
		c.JSON(http.StatusAccepted, gin.H{"success": true, "drawing_date": params.DateID, "winner": true, "lottery": params.Number, "result": result})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "drawing_date": params.DateID, "winner": false, "lottery": params.Number, "message": "โดนหวยแดก"})
}

func (h *Handler) FindLotteriesDate(c *gin.Context) {

}
