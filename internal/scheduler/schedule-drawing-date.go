package scheduler

import (
	"log"
	"lottery-web-scrapping/driver"
	drawingDateRepo "lottery-web-scrapping/internal/drawing-date/repository"
	drawingDateUsecase "lottery-web-scrapping/internal/drawing-date/usecase"
	drawingLotteryRepo "lottery-web-scrapping/internal/drawing-lottery/repository"
	drawingLotteryUsecase "lottery-web-scrapping/internal/drawing-lottery/usecase"
	webscrapper "lottery-web-scrapping/pkg/web-scrapper"
)

func ScheduleDrawingDate() {
	log.Println("################## Cron Job : Running Drawing Date #################")
	// Date
	drawingDateRepo := drawingDateRepo.NewDrawingDateRepository(driver.ClientMongo)
	drawingDateService := drawingDateUsecase.NewDrawingDateService(drawingDateRepo)

	lotterys, err := drawingDateService.CreateAllDate()
	if err != nil {
		log.Printf("Error Schedule : %v", err.Error())
	}

	// Lottery
	drawingLotteryRepo := drawingLotteryRepo.NewDrawingLotteryRepository(driver.ClientMongo)
	drawingLotteryService := drawingLotteryUsecase.NewDrawingLotteryService(drawingLotteryRepo)

	for _, lottery := range lotterys {
		data, err := webscrapper.GetLotteryByDate(lottery.Link)
		if err != nil {
			panic(err)
		}

		if err := drawingLotteryService.CreateDrawingLottery(data); err != nil {
			panic(err)
		}
	}

	log.Println("Schedule Drawing Date Completed")
}
