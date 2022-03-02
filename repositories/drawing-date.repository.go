package repositories

import (
	"context"
	"lottery-web-scrapping/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDrawingDate interface {
	CreateDrawingDate()
	FindByDate()
}

type DrawingDateService struct {
	BaseRepository
}

func NewDrawingDateService(c *mongo.Client) IDrawingDate {
	return &DrawingDateService{
		BaseRepository{c},
	}
}

func (c *DrawingDateService) CreateDrawingDate() {
	drawingDateCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("drawing_date")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}

func (c *DrawingDateService) FindByDate() {
	drawingDateCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("drawing_date")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
