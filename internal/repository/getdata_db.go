package repository

import (
	"context"
	"test/internal/models"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type GetDataDB struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
	cfg    *models.Config
}

func NewGetDataDB(db *sqlx.DB, logger *zap.SugaredLogger, cfg *models.Config) *GetDataDB {
	return &GetDataDB{
		db:     db,
		logger: logger,
		cfg:    cfg,
	}
}

func (r *GetDataDB) AddToDB(data *models.Data) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.RequestTimeout)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, "INSERT INTO links (active_link,history_link) VALUES ($1, $2)", data.Active_link, data.History_link)
	if err != nil {
		r.logger.Errorf("Error while insert data to DB: %v", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *GetDataDB) GetLinks(page int) (*[]models.Data, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.RequestTimeout)
	defer cancel()

	var links []models.Data
	var link models.Data

	row, err := r.db.QueryContext(ctx, "SELECT id, active_link, history_link FROM links ORDER BY id DESC LIMIT 30 OFFSET $1", page)
	if err != nil {
		r.logger.Errorf("Error while get links in DB: %v", err)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&link.Id, &link.Active_link, &link.History_link); err != nil {
			r.logger.Errorf("Error while scan dates in DB: %v", err)
			return nil, err
		}
		links = append(links, link)
	}
	return &links, nil
}

func (r *GetDataDB) GetLinkByID(id int) (*models.Data, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.RequestTimeout)
	defer cancel()

	var link models.Data
	if err := r.db.QueryRowContext(ctx, "SELECT id, active_link, history_link FROM links WHERE id=$1", id).Scan(&link.Id, &link.Active_link, &link.History_link); err != nil {
		r.logger.Errorf("Error while get links by ID in DB: %v", err)
		return &models.Data{}, err
	}
	return &link, nil
}

func (r *GetDataDB) AddLink(data models.Data) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.RequestTimeout)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query, err := r.db.PrepareContext(ctx, "INSERT INTO links(active_link, history_link) VALUES ($1, $2)")
	if err != nil {
		r.logger.Errorf("Error while add link in DB: %v", err)
		tx.Rollback()
		return err
	}
	if _, err := query.Exec(&data.Active_link, &data.History_link); err != nil {
		r.logger.Errorf("Error while add link in DB: %v", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *GetDataDB) UpdateLink(id int, data models.Data) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.RequestTimeout)
	defer cancel()

	link, err := r.GetLinkByID(id)
	if err != nil {
		r.logger.Errorf("Error while get link in DB:%v", err)
		return err
	}
	_, err = r.db.ExecContext(ctx, "UPDATE links SET active_link=$1,  history_link=$2 WHERE id=$3", data.Active_link, link.Active_link, id)
	if err != nil {
		r.logger.Errorf("Error while update link in DB: %v", err)
		return err
	}
	return nil
}

func (r *GetDataDB) DeleteLinkById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.RequestTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "DELETE FROM links WHERE id=$1", id)
	if err != nil {
		r.logger.Errorf("Error while delete link in DB: %v", err)
		return err
	}
	return nil
}
