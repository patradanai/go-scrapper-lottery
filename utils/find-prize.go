package utils

import (
	"lottery-web-scrapping/models"
)

type typePrize struct {
	Message string
	Name    string
	Prize   string
	Lottery string
}

func FindPrize(number string, data *models.DrawingLottery) ([]typePrize, bool) {
	var totalPrize []typePrize
	// Find First Prize
	if Contains(data.FirstPrize.Lottery, number) {
		lottery := &typePrize{Message: "Congratulations " + data.FirstPrize.Name, Name: data.FirstPrize.Name, Prize: data.FirstPrize.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	// Near First Prize
	if Contains(data.NearFirstPrize.Lottery, number) {
		lottery := &typePrize{Message: "Congratulations " + data.NearFirstPrize.Name, Name: data.NearFirstPrize.Name, Prize: data.NearFirstPrize.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	// Find Second Prize
	if Contains(data.SecondPrize.Lottery, number) {
		lottery := &typePrize{Message: "Congratulations " + data.SecondPrize.Name, Name: data.SecondPrize.Name, Prize: data.SecondPrize.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	// Find Third Prize
	if Contains(data.ThridPrize.Lottery, number) {
		lottery := &typePrize{Message: "Congratulations " + data.ThridPrize.Name, Name: data.ThridPrize.Name, Prize: data.ThridPrize.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	// Find Fourth Prize
	if Contains(data.FourthPrize.Lottery, number) {
		lottery := &typePrize{Message: "Congratulations " + data.FourthPrize.Name, Name: data.FourthPrize.Name, Prize: data.FourthPrize.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	// Find Fifth Prize
	if Contains(data.FifthPrize.Lottery, number) {
		lottery := &typePrize{Message: "Congratulations " + data.FifthPrize.Name, Name: data.FifthPrize.Name, Prize: data.FifthPrize.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	// Slice Number
	f3 := number[:3]
	l3 := number[3:6]
	l2 := number[4:]

	// Find Front Thrid Prize
	if Contains(data.FrontThird.Lottery, f3) {
		lottery := &typePrize{Message: "Congratulations " + data.FrontThird.Name, Name: data.FrontThird.Name, Prize: data.FrontThird.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	// Find End Thrid Prize
	if Contains(data.EndThird.Lottery, l3) {
		lottery := &typePrize{Message: "Congratulations " + data.EndThird.Name, Name: data.EndThird.Name, Prize: data.EndThird.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	// Find End Second Prize
	if Contains(data.EndSecond.Lottery, l2) {
		lottery := &typePrize{Message: "Congratulations " + data.EndSecond.Name, Name: data.EndSecond.Name, Prize: data.EndSecond.Prize, Lottery: number}
		totalPrize = append(totalPrize, *lottery)
	}

	exist := false

	if len(totalPrize) > 0 {
		exist = true
	}

	return totalPrize, exist
}
