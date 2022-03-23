package services

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IOAuthClientService interface {
	FindOAuthClient(apiKey string) (*models.OauthClient, error)
	FindOAuthByIDClient(userId string) (*models.OauthClient, error)
	CreateOAuthClient(userId string, apiKey string, clientId string, clientSecret string) error
	UpdateOAuthClient(userId string, oauthId string, apiKey string, clientId string, clientSecret string) error
}

type OAuthClientService struct {
	repository repositories.IOAuthClientRepository
}

func NewOAuthClientService(r repositories.IOAuthClientRepository) IOAuthClientService {
	return &OAuthClientService{r}
}

func (s *OAuthClientService) FindOAuthClient(apiKey string) (*models.OauthClient, error) {

	filters := bson.M{"api_key": apiKey}

	result, err := s.repository.FindOneOAuthClient(filters)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *OAuthClientService) FindOAuthByIDClient(userId string) (*models.OauthClient, error) {
	filter := bson.M{"user_id": userId}

	oauth, err := s.repository.FindOneOAuthClient(filter)
	if err != nil {
		return nil, err
	}

	return oauth, nil
}

func (s *OAuthClientService) CreateOAuthClient(userId string, apiKey string, clientId string, clientSecret string) error {
	objId, _ := primitive.ObjectIDFromHex(userId)
	oAuth := &models.OauthClient{
		UserId:       objId,
		ApiKey:       apiKey,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Revoke:       false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return s.repository.CreateOneOAuthClient(oAuth)
}

func (s *OAuthClientService) UpdateOAuthClient(userId string, oauthId string, apiKey string, clientId string, clientSecret string) error {
	filter := bson.M{"user_id": userId, "_id": oauthId}
	after := options.After
	opts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	oAuth := &models.OauthClient{
		ApiKey:       apiKey,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Revoke:       false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return s.repository.UpdateOne(filter, oAuth, opts)
}
