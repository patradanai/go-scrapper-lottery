package services

import (
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/role/repository"
)

type IRoleService interface {
	FindRoleByName(name string) (*models.Role, error)
}

type RoleService struct {
	repository repository.IRoleRepository
}

func NewRoleService(r repository.IRoleRepository) IRoleService {
	return &RoleService{r}
}

func (s *RoleService) FindRoleByName(name string) (*models.Role, error) {
	role, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return role, nil
}
