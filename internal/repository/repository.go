package repository

import (
	"test/internal/models"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type GetData interface {
	AddToDB(data *models.Data) error
	GetLinks(page int) (*[]models.Data, error)
	GetLinkByID(id int) (*models.Data, error)
	AddLink(data models.Data) error
	UpdateLink(id int, data models.Data) error
	DeleteLinkById(id int) error
}
type Client interface {
	Check(link string) (models.Data, error)
}

type Cache interface {
	Add(key, value string) error
	Get(key string) (string, bool, error)
}

type Repository struct {
	GetData
	Client
	Cache
}

func NewRepository(db *sqlx.DB, redis *redis.Client, logger *zap.SugaredLogger, cfg *models.Config) *Repository {
	return &Repository{
		GetData: NewGetDataDB(db, logger, cfg),
		Client:  NewClientDB(db, logger, cfg),
		Cache:   NewRedisCache(redis, logger),
	}
}
