package repositories

import (
	"context"
	"log"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRefreshTokenRepository interface {
	CreateOne(token *models.RefreshToken) error
	FindOne(token string, userId string) (*models.RefreshToken, error)
}

type RefreshTokenRepository struct {
	BaseRepository
}

func NewRefreshTokenRepository(c *mongo.Client) IRefreshTokenRepository {
	return &RefreshTokenRepository{
		BaseRepository{c},
	}
}

func (r *BaseRepository) CreateOne(token *models.RefreshToken) error {
	refreshTokenCollection := r.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("refresh_tokens")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := refreshTokenCollection.InsertOne(ctx, token); err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) FindOne(token string, userId string) (*models.RefreshToken, error) {
	refreshTokenCollection := r.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("refresh_tokens")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Printf("token :%v, userId: %v\n", token, userId)
	filter := bson.M{"token": token}
	refreshToken := models.RefreshToken{}

	if err := refreshTokenCollection.FindOne(ctx, filter).Decode(&refreshToken); err != nil {
		return nil, err
	}

	return &refreshToken, nil
}
