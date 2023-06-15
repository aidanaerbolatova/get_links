package repository

import (
	"context"
	"errors"
	"test/internal/models"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var ErrNoRows = errors.New("no rows found")

type ClientDB struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
	cfg    *models.Config
}

func NewClientDB(db *sqlx.DB, logger *zap.SugaredLogger, cfg *models.Config) *ClientDB {
	return &ClientDB{
		db:     db,
		logger: logger,
		cfg:    cfg,
	}
}

func (c *ClientDB) Check(link string) (models.Data, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.RequestTimeout)
	defer cancel()
	var data models.Data
	row := c.db.QueryRowContext(ctx, "SELECT id, active_link, history_link FROM links WHERE active_link=$1", link)
	err := row.Scan(&data.Id, &data.Active_link, &data.History_link)
	if err != nil {
		c.logger.Errorf("Error while check link in databse: %v", err)
		return models.Data{}, ErrNoRows
	}
	return data, nil
}
