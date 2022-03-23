package services

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/repositories"
)

type IRoleService interface {
	FindRoleByName(name string) (*models.Role, error)
}

type RoleService struct {
	repository repositories.IRoleRepository
}

func NewRoleService(r repositories.IRoleRepository) IRoleService {
	return &RoleService{r}
}

func (s *RoleService) FindRoleByName(name string) (*models.Role, error) {
	role, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return role, nil
}
