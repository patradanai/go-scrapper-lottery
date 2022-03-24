package middlewares

import (
	"lottery-web-scrapping/driver"
	oauthRepo "lottery-web-scrapping/internal/oauth/repository"
	oauthUsecase "lottery-web-scrapping/internal/oauth/usecase"
	httpError "lottery-web-scrapping/pkg/http-error"
	"lottery-web-scrapping/pkg/utils"

	"github.com/gin-gonic/gin"

	"strings"
)

type HeaderToken struct {
	Token string `header:"Authorization"`
}

type HeaderAPI struct {
	ApiKey string `header:"x-api-key" binding:"required"`
}

func AuthorizationJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenModel := HeaderToken{}
		if err := c.ShouldBindHeader(&tokenModel); err != nil {
			httpError.NewUnauthorize(c, "not exist token in header")
			c.Abort()
			return
		}

		token := strings.Split(tokenModel.Token, "Bearer ")
		if len(token) < 2 {
			httpError.NewBadRequest(c, "please check you token")
			c.Abort()
			return
		}

		jwtService := utils.NewJWTService()

		claims, err := jwtService.ValidateToken(token[1])
		if err != nil {
			httpError.NewInternalServerError(c, err.Error())
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)

		c.Next()
	}
}

func AuthorizationAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiModel := HeaderAPI{}

		if err := c.ShouldBindHeader(&apiModel); err != nil {
			httpError.NewBadRequest(c, "not exist token in header")
			c.Abort()
			return
		}

		oauthRepo := oauthRepo.NewOAuthClientRepository(driver.ClientMongo)
		oauthService := oauthUsecase.NewOAuthClientService(oauthRepo)

		result, err := oauthService.FindOAuthClient(apiModel.ApiKey)
		if err != nil {
			httpError.NewUnauthorize(c, "not exist api key")
			c.Abort()
			return
		}

		c.Set("userId", result.UserId)

		c.Next()
	}
}
