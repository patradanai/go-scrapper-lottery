package usecase

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/oauth/repository"
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
	repository repository.IOAuthClientRepository
}

func NewOAuthClientService(r repository.IOAuthClientRepository) IOAuthClientService {
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
	obj, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"user_id": obj}

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
	userObj, _ := primitive.ObjectIDFromHex(userId)
	oauthObj, _ := primitive.ObjectIDFromHex(oauthId)
	filter := bson.M{"user_id": userObj, "_id": oauthObj}
	after := options.After
	opts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	updates := bson.M{"$set": bson.M{"api_key": apiKey, "client_id": clientId, "client_secret": clientSecret, "updated_at": time.Now()}}

	return s.repository.UpdateOne(filter, updates, opts)
}
