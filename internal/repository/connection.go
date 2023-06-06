package repository

import (
	"fmt"
	"test/internal/models"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func ConnectDB(logger *zap.SugaredLogger, config *models.Config) (*sqlx.DB, error) {
	return NewPostgresDB(logger, config)
}

func NewPostgresDB(logger *zap.SugaredLogger, cfg *models.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logger.Errorf("Error while connect to Postgres: %v", err)
		return nil, err
	}
	return db, nil
}
