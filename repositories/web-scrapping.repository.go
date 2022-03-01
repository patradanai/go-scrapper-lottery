package repositories

import (
	"context"
	"lottery-web-scrapping/models"
	"lottery-web-scrapping/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IWebScrappingRepository interface {
	CreateLottery(lottery *models.DrawingLottery) error
	FindByNumber() (*models.DrawingLottery, error)
	FindByDate() (interface{}, error)
}

type WebScrappingRepository struct {
	BaseRepository
}

func NewWebScrappingRepository(c *mongo.Client) IWebScrappingRepository {
	return &WebScrappingRepository{
		BaseRepository{c},
	}
}

/*
	Create Lottery
*/
func (c *WebScrappingRepository) CreateLottery(lottery *models.DrawingLottery) error {
	lotteryCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("drawing_lottery")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := lotteryCollection.InsertOne(ctx, lottery); err != nil {
		return err
	}
	return nil
}

/*
	Find One Lottery by Filter Date , Number
*/
func (c *WebScrappingRepository) FindByNumber() (*models.DrawingLottery, error) {
	lotteryCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("drawing_lottery")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var lottery models.DrawingLottery
	filter := bson.M{}

	if err := lotteryCollection.FindOne(ctx, filter).Decode(&lottery); err != nil {
		return nil, err
	}

	return &lottery, nil
}

/*
	Find Many Lottery by Date
*/
func (c *WebScrappingRepository) FindByDate() (interface{}, error) {
	return nil, nil
}
