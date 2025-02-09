package service

import (
	"taskserver/dbserver/internal/entity"
	"taskserver/dbserver/internal/repository"
)

type ContainersService struct {
	repo repository.Containers
}

func NewContainersService(repo repository.Containers) *ContainersService {
	return &ContainersService{repo: repo}
}

func (s *ContainersService) Update(item entity.UpdateContainers) error {
	return s.repo.Update(item)
}

func (s *ContainersService) GetAll() ([]entity.Container, error) {
	return s.repo.GetAll()
}
