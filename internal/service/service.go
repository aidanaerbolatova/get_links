package service

import (
	"test/internal/models"
	"test/internal/repository"

	"go.uber.org/zap"
)

type GetData interface {
	AddToDB() ([]models.Data, error)
	GetLinks(page int) (*[]models.Data, error)
	GetLinkByID(id int) (*models.Data, error)
	AddLink(data models.Data) error
	UpdateLink(id int, data models.Data) error
	DeleteLinkById(id int) error
}

type Client interface {
	Check(link string) (int, error)
}

type Cache interface {
	Add(key, value string) error
	Get(key string) (string, bool, error)
}

type Service struct {
	GetData
	Client
	Cache
}

func NewService(repo *repository.Repository, logger *zap.SugaredLogger) *Service {
	return &Service{
		GetData: NewGetDataService(repo, logger),
		Client:  NewClientService(*repo),
		Cache:   NewCacheService(repository.RedisCache{}, logger),
	}
}
