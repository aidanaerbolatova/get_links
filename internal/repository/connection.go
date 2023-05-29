package repository

import (
	"test/config"

	"github.com/jmoiron/sqlx"
)

func ConnectDB() (*sqlx.DB, error) {
	var db *sqlx.DB
	config, err := config.ParseYaml()
	if err != nil {
		return db, err
	}
	return NewPostgresDB(config)
}
