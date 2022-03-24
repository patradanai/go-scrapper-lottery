package repository

import (
	"context"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IOAuthClientRepository interface {
	CreateOneOAuthClient(oauth *models.OauthClient) error
	UpdateOne(filter bson.M, update interface{}, opt *options.FindOneAndUpdateOptions) error
	FindOneOAuthClient(filter bson.M) (*models.OauthClient, error)
}

type OAuthClientRepository struct {
	*mongo.Client
}

func NewOAuthClientRepository(c *mongo.Client) IOAuthClientRepository {
	return &OAuthClientRepository{
		c,
	}
}

func (r *OAuthClientRepository) CreateOneOAuthClient(oauth *models.OauthClient) error {
	oauthClientCollection := r.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("oauth_clients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := oauthClientCollection.InsertOne(ctx, oauth); err != nil {
		return err
	}

	return nil
}

func (r *OAuthClientRepository) FindOneOAuthClient(filter bson.M) (*models.OauthClient, error) {
	oauthClientCollection := r.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("oauth_clients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oauthClient := models.OauthClient{}

	if err := oauthClientCollection.FindOne(ctx, filter).Decode(&oauthClient); err != nil {
		return nil, err
	}

	return &oauthClient, nil
}

func (r *OAuthClientRepository) UpdateOne(filter bson.M, update interface{}, opt *options.FindOneAndUpdateOptions) error {
	oauthClientCollection := r.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("oauth_clients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oauthClient := models.OauthClient{}

	if err := oauthClientCollection.FindOneAndUpdate(ctx, filter, update, opt).Decode(&oauthClient); err != nil {

		return err
	}

	return nil
}
