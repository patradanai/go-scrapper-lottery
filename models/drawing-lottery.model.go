package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LotteryType struct {
	Name    string
	Prize   string
	Lottery []int
}

type DrawingLottery struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FullDate       string             `bson:"full_date,omitempty"`
	Year           string             `bson:"year,omitempty"`
	Month          string             `bson:"month,omitempty"`
	Day            string             `bson:"day,omitempty"`
	FirstPrize     LotteryType        `bson:"first_prize"`
	FrontThird     LotteryType        `bson:"front_thrid"`
	EndThird       LotteryType        `bson:"end_thrid"`
	EndSecond      LotteryType        `bson:"end_second"`
	NearFirstPrize LotteryType        `bson:"near_first_prize"`
	SecondPrize    LotteryType        `bson:"second_prize"`
	ThridPrize     LotteryType        `bson:"third_prize"`
	FourthPrize    LotteryType        `bson:"fourth_prize"`
	FifthPrize     LotteryType        `bson:"fifth_prize"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}
