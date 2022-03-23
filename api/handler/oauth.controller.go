package handler

import (
	httperror "lottery-web-scrapping/pkg/http-error"
	"net/http"

	"github.com/gin-gonic/gin"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (h *Handler) CreateOAuth(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	apiKey, _ := gonanoid.New(20)
	clientId, _ := gonanoid.New(32)
	clientSecret, _ := gonanoid.New(32)

	if err := h.oauthService.CreateOAuthClient(userId, apiKey, clientId, clientSecret); err != nil {
		httperror.NewBadRequest(c, err.Error())
		return
	}

	data := make(map[string]string)
	data["apiKey"] = apiKey
	data["clientId"] = clientId
	data["clientSecret"] = clientSecret

	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func (h *Handler) UpdateOAuth(c *gin.Context) {
	oauthId := c.Param("oauthId")
	userId := c.MustGet("userId").(string)

	apiKey, _ := gonanoid.New(20)
	clientId, _ := gonanoid.New(32)
	clientSecret, _ := gonanoid.New(32)

	if err := h.oauthService.UpdateOAuthClient(userId, oauthId, apiKey, clientId, clientSecret); err != nil {
		httperror.NewBadRequest(c, err.Error())
		return
	}

	data := make(map[string]string)
	data["apiKey"] = apiKey
	data["clientId"] = clientId
	data["clientSecret"] = clientSecret

	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}
