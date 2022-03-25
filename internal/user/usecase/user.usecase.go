package usecase

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/user/repository"
)

type IUserService interface {
	CreateUser(user *models.User) error
	FindByUser(username string) (*models.User, error)
	FindById(id string) (*models.User, error)
}

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(r repository.IUserRepository) IUserService {
	return &UserService{r}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repository.CreateOne(user)
}

func (s *UserService) FindByUser(username string) (*models.User, error) {
	result, err := s.repository.FindUser(username)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *UserService) FindById(id string) (*models.User, error) {
	result, err := s.repository.FindId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
