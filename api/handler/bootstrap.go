package handler

import (
	"lottery-web-scrapping/internal/repositories"
	"lottery-web-scrapping/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	lotteryService      services.IDrawingLotteryService
	drawingDateService  services.IDrawingDateService
	userService         services.IUserService
	oauthService        services.IOAuthClientService
	refreshTokenService services.IRefreshTokenService
}

func NewHanlder(c *mongo.Client) *Handler {
	lotteryRepo := repositories.NewDrawingLotteryRepository(c)
	lotteryService := services.NewDrawingLotteryService(lotteryRepo)

	drawingDateRepo := repositories.NewDrawingDateRepository(c)
	drawingDateService := services.NewDrawingDateService(drawingDateRepo)

	userRepo := repositories.NewUserRepository(c)
	userService := services.NewUserService(userRepo)

	oauthRepo := repositories.NewOAuthClientRepository(c)
	oathService := services.NewOAuthClientService(oauthRepo)

	refreshTokenRepo := repositories.NewRefreshTokenRepository(c)
	refreshTokenService := services.NewRefreshTokenService(refreshTokenRepo)

	return &Handler{lotteryService: lotteryService, drawingDateService: drawingDateService, userService: userService, oauthService: oathService, refreshTokenService: refreshTokenService}
}
