package repository

import (
	"test/internal/models"

	"github.com/jmoiron/sqlx"
)

type GetData interface {
	AddToDB(data models.Data) error
	GetLinks(page int) ([]models.Data, error)
	GetLinkByID(id int) (models.Data, error)
	AddLink(data models.Data) error
	UpdateLink(id int, data models.Data) error
	DeleteLinkById(id int) error
}

type Repository struct {
	GetData
	Client
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		GetData: NewGetDataDB(db),
		Client:  NewClientDB(db),
	}
}
