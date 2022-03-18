package repositories

import (
	"context"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/internal/models"

	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDrawingDateRepository interface {
	CreateDrawingDate(date *models.DrawingDate) error
	FindByDate(date string) (*models.DrawingDate, error)
}

type DrawingDateRepository struct {
	BaseRepository
}

func NewDrawingDateRepository(c *mongo.Client) IDrawingDateRepository {
	return &DrawingDateRepository{
		BaseRepository{c},
	}
}

func (c *DrawingDateRepository) CreateDrawingDate(date *models.DrawingDate) error {
	drawingDateCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("drawing_date")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := drawingDateCollection.InsertOne(ctx, date); err != nil {
		return err
	}

	return nil
}

func (c *DrawingDateRepository) FindByDate(date string) (*models.DrawingDate, error) {
	drawingDateCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("drawing_date")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.M{"full_date": date}

	drawing := models.DrawingDate{}
	if err := drawingDateCollection.FindOne(ctx, filters).Decode(&drawing); err != nil {
		return nil, err
	}

	return &drawing, nil
}
