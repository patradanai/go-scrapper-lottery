package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"lottery-web-scrapping/models"
	"lottery-web-scrapping/utils"
	"time"
)

type IRefreshTokenRepository interface {
	CreateOneRefreshToken(token *models.RefreshToken) error
}

type RefreshTokenRepository struct {
	BaseRepository
}

func NewRefreshTokenRepository(c *mongo.Client) IRefreshTokenRepository {
	return &RefreshTokenRepository{
		BaseRepository{c},
	}
}

func (r *BaseRepository) CreateOneRefreshToken(token *models.RefreshToken) error {
	refreshTokenCollection := r.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("refresh_tokens")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := refreshTokenCollection.InsertOne(ctx, token); err != nil {
		return err
	}

	return nil
}
