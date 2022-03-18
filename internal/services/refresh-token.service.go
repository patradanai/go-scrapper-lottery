package services

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/repositories"
	"time"
)

type IRefreshTokenService interface {
	CreateRefreshToken(token string, expired int) error
}

type RefreshTokenService struct {
	repository repositories.IRefreshTokenRepository
}

func NewRefreshTokenService(r repositories.IRefreshTokenRepository) IRefreshTokenService {
	return &RefreshTokenService{r}
}

func (s *RefreshTokenService) CreateRefreshToken(token string, expired int) error {
	tokenModel := models.RefreshToken{}
	tokenModel.Token = token
	tokenModel.Revoke = false
	tokenModel.ExpiredAt = time.Now().AddDate(0, 0, expired)
	tokenModel.CreatedAt = time.Now()
	tokenModel.UpdatedAt = time.Now()

	if err := s.repository.CreateOneRefreshToken(&tokenModel); err != nil {
		return err
	}

	return nil
}
