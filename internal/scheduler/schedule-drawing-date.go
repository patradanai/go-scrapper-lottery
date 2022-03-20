package scheduler

import (
	"log"
	"lottery-web-scrapping/driver"
	"lottery-web-scrapping/internal/repositories"
	"lottery-web-scrapping/internal/services"
	webscrapper "lottery-web-scrapping/pkg/web-scrapper"
)

func ScheduleDrawingDate() {
	log.Println("################## Cron Job : Running Drawing Date #################")
	// Date
	drawingDateRepo := repositories.NewDrawingDateRepository(driver.ClientMongo)
	drawingDateService := services.NewDrawingDateService(drawingDateRepo)

	lotterys, err := drawingDateService.CreateAllDate()
	if err != nil {
		log.Printf("Error Schedule : %v", err.Error())
	}

	// Lottery
	drawingLotteryRepo := repositories.NewDrawingLotteryRepository(driver.ClientMongo)
	drawingLotteryService := services.NewDrawingLotteryService(drawingLotteryRepo)

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
