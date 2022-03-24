package handler

import (
	drawingDateRepo "lottery-web-scrapping/internal/drawing-date/repository"
	drawingDateUsecase "lottery-web-scrapping/internal/drawing-date/usecase"
	drawingLotteryRepo "lottery-web-scrapping/internal/drawing-lottery/repository"
	drawingLotteryUsecase "lottery-web-scrapping/internal/drawing-lottery/usecase"
	oauthRepo "lottery-web-scrapping/internal/oauth/repository"
	oauthUsecase "lottery-web-scrapping/internal/oauth/usecase"
	refreshRepo "lottery-web-scrapping/internal/refresh-token/repository"
	refreshUsecase "lottery-web-scrapping/internal/refresh-token/usecase"
	roleRepo "lottery-web-scrapping/internal/role/repository"
	roleUsecase "lottery-web-scrapping/internal/role/usecase"
	userRepo "lottery-web-scrapping/internal/user/repository"
	userUsecase "lottery-web-scrapping/internal/user/usecase"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	lotteryService      drawingLotteryUsecase.IDrawingLotteryService
	drawingDateService  drawingDateUsecase.IDrawingDateService
	userService         userUsecase.IUserService
	oauthService        oauthUsecase.IOAuthClientService
	refreshTokenService refreshUsecase.IRefreshTokenService
	roleService         roleUsecase.IRoleService
}

func NewHanlder(c *mongo.Client) *Handler {
	lotteryRepo := drawingLotteryRepo.NewDrawingLotteryRepository(c)
	lotteryService := drawingLotteryUsecase.NewDrawingLotteryService(lotteryRepo)

	drawingDateRepo := drawingDateRepo.NewDrawingDateRepository(c)
	drawingDateService := drawingDateUsecase.NewDrawingDateService(drawingDateRepo)

	userRepo := userRepo.NewUserRepository(c)
	userService := userUsecase.NewUserService(userRepo)

	oauthRepo := oauthRepo.NewOAuthClientRepository(c)
	oathService := oauthUsecase.NewOAuthClientService(oauthRepo)

	refreshTokenRepo := refreshRepo.NewRefreshTokenRepository(c)
	refreshTokenService := refreshUsecase.NewRefreshTokenService(refreshTokenRepo)

	roleRepo := roleRepo.NewRoleRepository(c)
	roleService := roleUsecase.NewRoleService(roleRepo)

	return &Handler{lotteryService: lotteryService, drawingDateService: drawingDateService, userService: userService, roleService: roleService, oauthService: oathService, refreshTokenService: refreshTokenService}
}
