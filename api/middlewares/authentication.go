package middlewares

import (
	"lottery-web-scrapping/driver"
	refreshRepo "lottery-web-scrapping/internal/refresh-token/repository"
	refreshUsecase "lottery-web-scrapping/internal/refresh-token/usecase"
	userRepo "lottery-web-scrapping/internal/user/repository"
	userUsecase "lottery-web-scrapping/internal/user/usecase"
	httpError "lottery-web-scrapping/pkg/http-error"
	"lottery-web-scrapping/pkg/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Binding JSON
		userLogin := UserLogin{}
		if err := c.ShouldBindJSON(&userLogin); err != nil {
			httpError.NewBadRequest(c, nil)
			return
		}

		// Find User
		userRepo := userRepo.NewUserRepository(driver.ClientMongo)
		userService := userUsecase.NewUserService(userRepo)

		user, exist := userService.FindByUser(userLogin.Username)
		if !exist {
			httpError.NewRequestNotFound(c, nil)
			return
		}

		// Check Password
		if !utils.DecryptPwd(user.Password, userLogin.Password) {
			httpError.NewBadRequest(c, "user or password invalid")
			return
		}

		// Jwt
		jwtService := utils.NewJWTService()
		token, err := jwtService.GenerateToken(user.ID.String())
		if err != nil {
			httpError.NewInternalServerError(c, err.Error())
			return
		}

		// RefreshToken
		tokenRefresh := utils.GenUUID()
		refreshRepo := refreshRepo.NewRefreshTokenRepository(driver.ClientMongo)
		refreshService := refreshUsecase.NewRefreshTokenService(refreshRepo)
		if err := refreshService.CreateRefreshToken(tokenRefresh, 15); err != nil {
			httpError.NewInternalServerError(c, err.Error())
			return
		}

		// Roles

		// Map
		data := make(map[string]string)
		data["access_token"] = token
		data["refresh_token"] = tokenRefresh
		data["roles"] = ""

		c.JSON(http.StatusOK, gin.H{"success": true, "data": data})

	}
}
