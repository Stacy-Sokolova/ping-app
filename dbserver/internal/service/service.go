package service

import (
	"taskserver/dbserver/internal/entity"
	"taskserver/dbserver/internal/repository"
)

type Containers interface {
	GetAll() ([]entity.Container, error)
	Update(item entity.UpdateContainers) error
}

type Service struct {
	Containers
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Containers: NewContainersService(repos.Containers),
	}
}
