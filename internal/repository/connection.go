package repository

import (
	"context"
	"fmt"
	"test/internal/models"
	"time"

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

	CreateTables(db)

	return db, nil
}
func CreateTables(db *sqlx.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	query := `
	CREATE TABLE IF NOT EXISTS links(
		id SERIAL PRIMARY KEY,
		active_link VARCHAR(512),
		history_link VARCHAR(512)
	)
	`
	db.MustExecContext(ctx, query)
}
