package service

import (
	"test/internal/models"
	"test/internal/repository"
)

type GetData interface {
	AddToDB(data []models.Data) error
	GetLinks(page int) ([]models.Data, error)
	GetLinkByID(id int) (models.Data, error)
	AddLink(data models.Data) error
	UpdateLink(id int, data models.Data) error
	DeleteLinkById(id int) error
}

type Service struct {
	GetData
	Cache
	Client
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		GetData: NewGetDataService(repo),
		// Cache:   NewCacheService(),
	}
}
