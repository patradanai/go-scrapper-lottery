package services

import (
	"fmt"
	"lottery-web-scrapping/repositories"
	"lottery-web-scrapping/utils"
)

type IDrawingLotteryService interface {
	FindLotteryByNumber(number string, date string) (interface{}, bool)
	FindLotteryByDate(date string) (interface{}, error)
}

type DrawingLotteryService struct {
	repository repositories.IDrawingLotteryRepository
}

func NewDrawingLotteryService(r repositories.IDrawingLotteryRepository) IDrawingLotteryService {
	return &DrawingLotteryService{repository: r}
}

func (r *DrawingLotteryService) FindLotteryByNumber(number string, date string) (interface{}, bool) {
	result, err := r.repository.FindByDate(date)
	if err != nil {
		panic(err)
	}

	// Find Prize
	fmt.Println("Find Prize")
	data, exist := utils.FindPrize(number, result)

	return data, exist
}

func (r *DrawingLotteryService) FindLotteryByDate(date string) (interface{}, error) {
	result, err := r.repository.FindByDate(date)
	if err != nil {
		return nil, err
	}

	return result, nil
}
