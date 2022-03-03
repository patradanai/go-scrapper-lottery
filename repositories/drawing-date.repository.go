package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"lottery-web-scrapping/models"
	"lottery-web-scrapping/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDrawingDate interface {
	CreateDrawingDate(date *models.DrawingDate) error
	FindByDate(date string) (*models.DrawingDate, error)
}

type DrawingDateService struct {
	BaseRepository
}

func NewDrawingDateService(c *mongo.Client) IDrawingDate {
	return &DrawingDateService{
		BaseRepository{c},
	}
}

func (c *DrawingDateService) CreateDrawingDate(date *models.DrawingDate) error {
	drawingDateCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("drawing_date")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := drawingDateCollection.InsertOne(ctx, date); err != nil {
		return err
	}

	return nil
}

func (c *DrawingDateService) FindByDate(date string) (*models.DrawingDate, error) {
	drawingDateCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("drawing_date")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{"full_date": date}

	drawing := models.DrawingDate{}
	if err := drawingDateCollection.FindOne(ctx, filters).Decode(&drawing); err != nil {
		return nil, err
	}

	return &drawing, nil
}
