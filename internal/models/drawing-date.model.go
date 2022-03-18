package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LotteryLink struct {
	Link string
	Date string
}

type DrawingDate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	DateThai    string             `bson:"date_th,omitempty"`
	DateEng     string             `bson:"date_en,omitempty"`
	FullDate    string             `bson:"full_date,omitempty"`
	Year        string             `bson:"year,omitempty"`
	Month       string             `bson:"month,omitempty"`
	Day         string             `bson:"day,omitempty"`
	LotteryLink LotteryLink        `bson:"lottery_link"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
