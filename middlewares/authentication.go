package middlewares

import (
	"github.com/gin-gonic/gin"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/repositories"
	"lottery-web-scrapping/services"
	"lottery-web-scrapping/utils"
	"net/http"
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
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Please Check your body"})
			return
		}

		// Find User
		userRepo := repositories.NewUserRepository(configs.ClientMongo)
		userService := services.NewUserService(userRepo)

		user, err := userService.FindByUser(userLogin.Username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "user not exist in system"})
			return
		}

		// Check Password
		if utils.DecryptPwd(user.Password, userLogin.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "user or password invalid"})
			return
		}

		// Jwt
		jwtService := utils.NewJWTService()
		token, err := jwtService.GenerateToken(user.ID.String())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "something went wrong"})
			return
		}

		// RefreshToken
		tokenRefresh := utils.GenUUID()
		refreshRepo := repositories.NewRefreshTokenRepository(configs.ClientMongo)
		refreshService := services.NewRefreshTokenService(refreshRepo)
		if err := refreshService.CreateRefreshToken(tokenRefresh, 15); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "something went wrong"})
			return
		}

		// Roles

		// Map
		data := make(map[string]string)
		data["access_token"] = token
		data["refresh_token"] = tokenRefresh
		data["roles"] = ""

		c.JSON(http.StatusOK, gin.H{"success": true})

	}
}
