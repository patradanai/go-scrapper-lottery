package utils

import (
	"fmt"
	"lottery-web-scrapping/models"
)

func FindPrize(number string, data *models.DrawingLottery) {
	// Find First Prize
	if Contains(data.FirstPrize.Lottery, number) {
		fmt.Println(data.FirstPrize.Name)
	}

	// Near First Prize
	if Contains(data.NearFirstPrize.Lottery, number) {
		fmt.Println(data.NearFirstPrize.Name)
	}

	// Find Second Prize
	if Contains(data.SecondPrize.Lottery, number) {
		fmt.Println(data.SecondPrize.Name)
	}

	// Find Third Prize
	if Contains(data.ThridPrize.Lottery, number) {
		fmt.Println(data.ThridPrize.Name)
	}

	// Find Fourth Prize
	if Contains(data.FourthPrize.Lottery, number) {
		fmt.Println(data.FourthPrize.Name)
	}

	// Find Fifth Prize
	if Contains(data.FifthPrize.Lottery, number) {
		fmt.Println(data.FifthPrize.Name)
	}

	// Slice Number
	f3 := number[:3]
	l3 := number[3:6]
	l2 := number[4:]

	// Find Front Thrid Prize
	if Contains(data.FrontThird.Lottery, f3) {
		fmt.Println(data.FrontThird.Name)
	}

	// Find End Thrid Prize
	if Contains(data.EndThird.Lottery, l3) {
		fmt.Println(data.EndThird.Name)
	}

	// Find End Second Prize
	if Contains(data.EndSecond.Lottery, l2) {
		fmt.Println(data.EndSecond.Name)
	}
}
