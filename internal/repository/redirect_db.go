package repository

import (
	"test/internal/models"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Client interface {
	Check(link string) (models.Data, error)
}

type ClientDB struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewClientDB(db *sqlx.DB, logger *zap.SugaredLogger) *ClientDB {
	return &ClientDB{
		db:     db,
		logger: logger,
	}
}

func (c *ClientDB) Check(link string) (models.Data, error) {
	var data models.Data
	row := c.db.QueryRow("SELECT id, active_link, history_link FROM links WHERE active_link=$1", link)
	err := row.Scan(&data.Id, &data.Active_link, &data.History_link)
	if err != nil {
		c.logger.Errorf("Error while check link in databse: %v", err)
		return models.Data{}, err
	}
	return data, nil
}
