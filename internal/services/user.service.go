package services

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/repositories"
)

type IUserService interface {
	FindByUser(username string) (*models.User, error)
	FindById(id string) (*models.User, error)
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(r repositories.IUserRepository) IUserService {
	return &UserService{r}
}

func (s *UserService) FindByUser(username string) (*models.User, error) {
	result, err := s.FindByUser(username)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) FindById(id string) (*models.User, error) {
	result, err := s.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
