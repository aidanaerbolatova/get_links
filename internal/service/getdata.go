package service

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"test/internal/models"
	"test/internal/repository"
)

type GetDataService struct {
	repo repository.GetData
}

func NewGetDataService(repo repository.GetData) *GetDataService {
	return &GetDataService{repo: repo}
}

func ConvertJson(file string) ([]models.Data, error) {
	var data []models.Data
	jsonfile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *GetDataService) AddToDB(data []models.Data) error {
	for _, link := range data {
		if err := s.repo.AddToDB(link); err != nil {
			return err
		}
	}
	return nil
}

func (s *GetDataService) GetLinks(page int) ([]models.Data, error) {
	if page != 1 {
		page = (page - 1) * 30
	} else {
		page -= 1
	}
	links, err := s.repo.GetLinks(page)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (s *GetDataService) GetLinkByID(id int) (models.Data, error) {
	link, err := s.repo.GetLinkByID(id)
	if err != nil {
		return models.Data{}, err
	}
	return link, nil
}

func (s *GetDataService) AddLink(data models.Data) error {
	return s.repo.AddLink(data)
}

func (s *GetDataService) UpdateLink(id int, data models.Data) error {
	return s.repo.UpdateLink(id, data)
}

func (s *GetDataService) DeleteLinkById(id int) error {
	return s.repo.DeleteLinkById(id)
}
