package repository

import (
	"test/internal/models"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func ConnectDB(logger *zap.SugaredLogger, config *models.Config) (*sqlx.DB, error) {
	return NewPostgresDB(logger, config)
}
