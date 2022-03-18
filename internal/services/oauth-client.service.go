package services

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
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
