package repository

import (
	"test/internal/models"

	"github.com/jmoiron/sqlx"
)

type GetDataDB struct {
	db *sqlx.DB
}

func NewGetDataDB(db *sqlx.DB) *GetDataDB {
	return &GetDataDB{db: db}
}

func (r *GetDataDB) AddToDB(data models.Data) error {
	_, err := r.db.Exec("INSERT INTO links (active_link,history_link) VALUES ($1, $2)", data.Active_link, data.History_link)
	if err != nil {
		return err
	}
	return nil
}

func (r *GetDataDB) GetLinks(page int) ([]models.Data, error) {
	var links []models.Data
	var link models.Data
	row, err := r.db.Query("SELECT id, active_link, history_link FROM links ORDER BY id DESC LIMIT 30 OFFSET $1", page)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&link.Id, &link.Active_link, &link.History_link); err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}

func (r *GetDataDB) GetLinkByID(id int) (models.Data, error) {
	var link models.Data
	if err := r.db.QueryRow("SELECT id, active_link, history_link FROM links WHERE id=$1", id).Scan(&link.Id, &link.Active_link, &link.History_link); err != nil {
		return models.Data{}, err
	}
	return link, nil
}

func (r *GetDataDB) AddLink(data models.Data) error {
	query, err := r.db.Prepare("INSERT INTO links(active_link, history_link) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	if _, err := query.Exec(data.Active_link, data.History_link); err != nil {
		return err
	}
	return nil
}

func (r *GetDataDB) UpdateLink(id int, data models.Data) error {
	link, err := r.GetLinkByID(id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("UPDATE links SET active_link=$1,  history_link=$2 WHERE id=$3", data.Active_link, link.Active_link, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *GetDataDB) DeleteLinkById(id int) error {
	_, err := r.db.Exec("DELETE FROM links WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
