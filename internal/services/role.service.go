package services

import "lottery-web-scrapping/internal/repositories"

type IRoleService interface {
}

type RoleService struct {
	repository repositories.IRoleRepository
}

func NewRoleService(r repositories.IRoleRepository) IRoleService {
	return &RoleService{r}
}
