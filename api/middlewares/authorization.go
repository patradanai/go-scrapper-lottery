package middlewares

import (
	"lottery-web-scrapping/driver"
	"lottery-web-scrapping/internal/repositories"
	"lottery-web-scrapping/internal/services"
	"lottery-web-scrapping/pkg/utils"

	"github.com/gin-gonic/gin"

	"net/http"
	"strings"
)

type HeaderToken struct {
	Token string `header:"Authorization"`
}

type HeaderAPI struct {
	ApiKey string `header:"x-api-key"`
}

func AuthorizationJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenModel := HeaderToken{}
		if err := c.ShouldBindHeader(&tokenModel); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "not exist token in header"})
			return
		}

		token := strings.Split(tokenModel.Token, "Bearer ")

		jwtService := utils.NewJWTService()

		claims, err := jwtService.ValidateToken(token[0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "something went wrong"})
			return
		}

		c.Set("userId", claims.Id)

		c.Next()
	}
}

func AuthorizationAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiModel := HeaderAPI{}

		if err := c.ShouldBindHeader(&apiModel); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "not exist token in header"})
			return
		}

		oauthRepo := repositories.NewOAuthClientRepository(driver.ClientMongo)
		oauthService := services.NewOAuthClientService(oauthRepo)

		result, err := oauthService.FindOAuthClient(apiModel.ApiKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "not exist api key"})
			return
		}

		c.Set("userId", result.UserId)

		c.Next()
	}
}
