package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
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
