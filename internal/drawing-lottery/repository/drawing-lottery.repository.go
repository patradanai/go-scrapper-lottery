package repository

import (
	"context"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDrawingLotteryRepository interface {
	CreateLottery(lottery *models.DrawingLottery) error
	FindByDate(date string) (*models.DrawingLottery, error)
}

type DrawingLotteryRepository struct {
	*mongo.Client
}

func NewDrawingLotteryRepository(c *mongo.Client) IDrawingLotteryRepository {
	return &DrawingLotteryRepository{
		c,
	}
}

/*
	Create Lottery
*/
func (c *DrawingLotteryRepository) CreateLottery(lottery *models.DrawingLottery) error {
	lotteryCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("drawing_lottery")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := lotteryCollection.InsertOne(ctx, lottery); err != nil {
		return err
	}
	return nil
}

/*
	Find Many Lottery by Date
*/
func (c *DrawingLotteryRepository) FindByDate(date string) (*models.DrawingLottery, error) {
	lotteryCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("drawing_lottery")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var lottery models.DrawingLottery

	filter := bson.M{"full_date": date}

	if err := lotteryCollection.FindOne(ctx, filter).Decode(&lottery); err != nil {
		return nil, err
	}

	return &lottery, nil
}
