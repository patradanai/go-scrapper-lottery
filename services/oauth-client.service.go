package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"lottery-web-scrapping/models"
	"lottery-web-scrapping/repositories"
)

type IOAuthClientService interface {
	FindOAuthClient(apiKey string) (*models.OauthClient, error)
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
