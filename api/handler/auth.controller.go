package handler

import (
	httpError "lottery-web-scrapping/pkg/http-error"

	"github.com/gin-gonic/gin"
)

type registerBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	login := registerBody{}

	if err := c.ShouldBindJSON(&login); err != nil {
		httpError.NewBadRequest(c, "require field username and password")
		return
	}

}
