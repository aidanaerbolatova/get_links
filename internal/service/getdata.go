package service

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"test/internal/models"
	"test/internal/repository"

	"go.uber.org/zap"
)

type GetDataService struct {
	repo   repository.GetData
	logger *zap.SugaredLogger
}

func NewGetDataService(repo repository.GetData, logger *zap.SugaredLogger) *GetDataService {
	return &GetDataService{repo: repo, logger: logger}
}

func ConvertJson(file string, logger *zap.SugaredLogger) ([]models.Data, error) {
	var data []models.Data
	jsonfile, err := os.Open(file)
	if err != nil {
		logger.Errorf("Error while open file: %v", err)
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		logger.Errorf("Error while read file: %v", err)
		return nil, err
	}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		logger.Errorf("Error while unmarshal file: %v", err)
		return nil, err
	}
	return data, nil
}

func (s *GetDataService) AddToDB() error {
	data, err := ConvertJson("links.json", s.logger)
	if err != nil {
		s.logger.Errorf("Error while convert json: %v", err)
		return err
	}
	for _, link := range data {
		if err := s.repo.AddToDB(&link); err != nil {
			s.logger.Errorf("Error while add to DB: %v", err)
			return err
		}
	}
	return nil
}

func (s *GetDataService) GetLinks(page int) (*[]models.Data, error) {
	if page != 1 {
		page = (page - 1) * 30
	} else {
		page -= 1
	}
	links, err := s.repo.GetLinks(page)
	if err != nil {
		s.logger.Errorf("Error while get links: %v", err)
		return nil, err
	}
	return links, nil
}

func (s *GetDataService) GetLinkByID(id int) (*models.Data, error) {
	link, err := s.repo.GetLinkByID(id)
	if err != nil {
		s.logger.Errorf("Error while get link by ID: %v", err)
		return &models.Data{}, err
	}
	return link, nil
}

func (s *GetDataService) AddLink(data models.Data) error {
	err := s.repo.AddLink(data)
	if err != nil {
		s.logger.Errorf("Error while add link: %v", err)
		return err
	}
	return nil
}

func (s *GetDataService) UpdateLink(id int, data models.Data) error {
	err := s.repo.UpdateLink(id, data)
	if err != nil {
		s.logger.Errorf("Error while update link: %v", err)
		return err
	}
	return nil
}

func (s *GetDataService) DeleteLinkById(id int) error {
	err := s.repo.DeleteLinkById(id)
	if err != nil {
		s.logger.Errorf("Error while delete link: %v", err)
		return err
	}
	return nil
}
