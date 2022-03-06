package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"lottery-web-scrapping/models"
	"lottery-web-scrapping/utils"
	"time"
)

type IOAuthClientRepository interface {
	CreateOneOAuthClient(oauth *models.OauthClient) error
	FindOneOAuthClient(filter bson.M) (*models.OauthClient, error)
}

type OAuthClientRepository struct {
	BaseRepository
}

func NewOAuthClientRepository(c *mongo.Client) IOAuthClientRepository {
	return &OAuthClientRepository{
		BaseRepository{c},
	}
}

func (c *BaseRepository) CreateOneOAuthClient(oauth *models.OauthClient) error {
	oauthClientCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("oauth_clients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := oauthClientCollection.InsertOne(ctx, oauth); err != nil {
		return err
	}

	return nil
}

func (c *BaseRepository) FindOneOAuthClient(filter bson.M) (*models.OauthClient, error) {
	oauthClientCollection := c.Client.Database(utils.LoadEnv("MONGO_DB_NAME")).Collection("oauth_clients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oauthClient := models.OauthClient{}

	if err := oauthClientCollection.FindOne(ctx, filter).Decode(&oauthClient); err != nil {
		return nil, err
	}

	return &oauthClient, nil
}
