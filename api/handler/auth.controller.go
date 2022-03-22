package handler

import (
	"lottery-web-scrapping/internal/models"
	httpError "lottery-web-scrapping/pkg/http-error"
	"lottery-web-scrapping/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *Handler) SignUp(c *gin.Context) {
	login := registerBody{}

	if err := c.ShouldBindJSON(&login); err != nil {
		httpError.NewBadRequest(c, "require field username and password")
		return
	}

	_, exist := h.UserService.FindByUser(login.Username)
	if exist {
		httpError.NewBadRequest(c, "existing username")
		return
	}

	// Decode Password
	hashPwd, _ := utils.EncryptPwd(login.Password)

	user := &models.User{
		Username: login.Email,
		Password: hashPwd,
		Active:   true,
	}

	if err := h.UserService.CreateUser(user); err != nil {
		httpError.NewInternalServerError(c, nil)
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
