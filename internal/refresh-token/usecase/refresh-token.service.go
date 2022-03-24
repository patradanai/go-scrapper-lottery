package usecase

import (
	"fmt"
	"log"
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/refresh-token/repository"
	"time"
)

type IRefreshTokenService interface {
	CreateRefreshToken(token string, expired int) error
	ValidateRefreshToken(token string) (*models.RefreshToken, error)
}

type RefreshTokenService struct {
	repository repository.IRefreshTokenRepository
}

func NewRefreshTokenService(r repository.IRefreshTokenRepository) IRefreshTokenService {
	return &RefreshTokenService{r}
}

func (s *RefreshTokenService) CreateRefreshToken(token string, expired int) error {
	tokenModel := models.RefreshToken{}
	tokenModel.Token = token
	tokenModel.Revoke = false
	tokenModel.ExpiredAt = time.Now().AddDate(0, 0, expired)
	tokenModel.CreatedAt = time.Now()
	tokenModel.UpdatedAt = time.Now()

	if err := s.repository.CreateOne(&tokenModel); err != nil {
		return err
	}

	return nil
}

func (s *RefreshTokenService) ValidateRefreshToken(token string) (*models.RefreshToken, error) {
	refreshToken, err := s.repository.FindOne(token)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	current := time.Now()
	expired := refreshToken.ExpiredAt
	log.Println(current, expired)
	if current.After(expired) {
		err := fmt.Errorf("token : %v expired", token)

		return nil, err
	}

	return refreshToken, nil
}
