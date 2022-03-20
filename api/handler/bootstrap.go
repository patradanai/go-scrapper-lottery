package handler

import (
	"lottery-web-scrapping/internal/repositories"
	"lottery-web-scrapping/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	LotteryService     services.IDrawingLotteryService
	DrawingDateService services.IDrawingDateService
	UserService        services.IUserService
}

func NewHanlder(c *mongo.Client) *Handler {
	lotteryRepo := repositories.NewDrawingLotteryRepository(c)
	lotteryService := services.NewDrawingLotteryService(lotteryRepo)

	drawingDateRepo := repositories.NewDrawingDateRepository(c)
	drawingDateService := services.NewDrawingDateService(drawingDateRepo)

	userRepo := repositories.NewUserRepository(c)
	userService := services.NewUserService(userRepo)

	return &Handler{LotteryService: lotteryService, DrawingDateService: drawingDateService, UserService: userService}
}
