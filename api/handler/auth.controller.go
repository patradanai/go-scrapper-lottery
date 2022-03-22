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

	_, exist := h.userService.FindByUser(login.Username)
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

	if err := h.userService.CreateUser(user); err != nil {
		httpError.NewInternalServerError(c, nil)
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

type refreshTokenBody struct {
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}

func (h *Handler) RefreshToken(c *gin.Context) {
	refreshToken := refreshTokenBody{}

	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		httpError.NewBadRequest(c, "refresh token must exist in body")
		return
	}

	// Valdate RefreshToken
	result, err := h.refreshTokenService.ValidateRefreshToken(refreshToken.RefreshToken, refreshToken.UserID)
	if err != nil {
		httpError.NewBadRequest(c, "refresh token is expired")
		return
	}

	jwtService := utils.NewJWTService()
	token, err := jwtService.GenerateToken(result.UserID.String())
	if err != nil {
		httpError.NewInternalServerError(c, nil)
		return
	}

	data := make(map[string]string)
	data["access_token"] = token
	data["refresh_token"] = refreshToken.RefreshToken

	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})

}

func (h *Handler) RevokeToken(c *gin.Context) {

}
