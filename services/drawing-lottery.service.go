package services

import (
	"fmt"
	"lottery-web-scrapping/repositories"
	"lottery-web-scrapping/utils"
)

type IDrawingLotteryService interface {
	FindLotteryByNumber(number string)
}

type DrawingLotteryService struct {
	repository repositories.IDrawingLotteryRepository
}

func NewDrawingLotteryService(r repositories.IDrawingLotteryRepository) IDrawingLotteryService {
	return &DrawingLotteryService{repository: r}
}

func (r *DrawingLotteryService) FindLotteryByNumber(number string) {
	result, err := r.repository.FindByNumber(number)
	if err != nil {
		panic(err)
	}

	// Find Prize
	fmt.Println("Find Prize")
	utils.FindPrize("360307", result)

	// jsonData, err := json.Marshal(result)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(jsonData))
}
