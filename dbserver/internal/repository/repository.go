package repository

import (
	"taskserver/dbserver/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Containers interface {
	GetAll() ([]entity.Container, error)
	Update(item entity.UpdateContainers) error
}

type Repository struct {
	Containers
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Containers: NewContainersPostgres(db),
	}
}
