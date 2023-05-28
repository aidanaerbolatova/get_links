package repository

import (
	"test/internal/models"

	"github.com/jmoiron/sqlx"
)

type Client interface {
	Check(link string) (models.Data, error)
}

type ClientDB struct {
	db *sqlx.DB
}

func NewClientDB(db *sqlx.DB) *ClientDB {
	return &ClientDB{
		db: db,
	}
}

func (c *ClientDB) Check(link string) (models.Data, error) {
	var data models.Data
	row := c.db.QueryRow("SELECT id, active_link, history_link FROM links WHERE active_link=$1", link)
	err := row.Scan(&data.Id, &data.Active_link, &data.History_link)
	if err != nil {
		return models.Data{}, err
	}
	return data, nil
}
