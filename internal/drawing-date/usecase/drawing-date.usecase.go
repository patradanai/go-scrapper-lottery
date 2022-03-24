package usecase

import (
	"fmt"
	"lottery-web-scrapping/internal/drawing-date/repository"
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/pkg/utils"
	webscrapper "lottery-web-scrapping/pkg/web-scrapper"
	"time"
)

type IDrawingDateService interface {
	CreateAllDate() ([]models.LotteryLink, error)
	FindDrawingDate(date string) error
}

type DrawingDateService struct {
	repository repository.IDrawingDateRepository
}

func NewDrawingDateService(r repository.IDrawingDateRepository) IDrawingDateService {
	return &DrawingDateService{repository: r}
}

func (r *DrawingDateService) CreateAllDate() ([]models.LotteryLink, error) {

	paramLink, err := webscrapper.FindAllDrawingDate()
	if err != nil {
		return nil, err
	}

	addDrawing := make([]models.LotteryLink, 0)

	for _, dateEn := range paramLink {

		date := utils.GetDate(dateEn.Date)

		dayString := fmt.Sprintf("%02d", utils.ConvToInteger(date["Day"]))
		monthlyString := fmt.Sprintf("%02d", utils.ConvMonthlyToNum(date["Monthly"]))
		yearString := fmt.Sprintf("%04d", utils.ConvToInteger(date["Year"]))

		fullDate := fmt.Sprintf("%v%v%v", dayString, monthlyString, yearString)

		_, err := r.repository.FindByDate(fullDate)
		if err == nil {
			continue
		}

		// Add in Slice
		addDrawing = append(addDrawing, dateEn)

		// Create Date
		lotteryDate := &models.DrawingDate{
			DateThai:    date["Monthly"],
			FullDate:    fullDate,
			Year:        yearString,
			Month:       monthlyString,
			Day:         dayString,
			LotteryLink: dateEn,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := r.repository.CreateDrawingDate(lotteryDate); err != nil {
			panic(err)
		}
	}

	return addDrawing, nil

}

func (r *DrawingDateService) FindDrawingDate(date string) error {
	if _, err := r.repository.FindByDate(date); err != nil {
		return err
	}

	return nil
}
