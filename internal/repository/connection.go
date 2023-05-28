package repository

import (
	"github.com/jmoiron/sqlx"
)

func ConnectDB() (*sqlx.DB, error) {
	return NewPostgresDB(Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
}
