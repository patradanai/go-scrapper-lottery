package services

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	CreateUser(user *models.User) error
	FindByUser(username string) (*models.User, bool)
	FindById(id string) (*models.User, error)
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(r repositories.IUserRepository) IUserService {
	return &UserService{r}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repository.CreateOne(user)
}

func (s *UserService) FindByUser(username string) (*models.User, bool) {
	result, err := s.repository.FindUser(username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false
		}
		return nil, true
	}
	return result, false
}

func (s *UserService) FindById(id string) (*models.User, error) {
	result, err := s.repository.FindId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
