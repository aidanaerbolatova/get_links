package repository

import (
	"test/internal/models"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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

func NewRepository(db *sqlx.DB, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		GetData: NewGetDataDB(db, logger),
		Client:  NewClientDB(db, logger),
	}
}
