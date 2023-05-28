package service

import (
	"database/sql"
	"errors"
	"test/internal/repository"
)

type Client interface {
	Check(link string) (int, error)
}

type ClientService struct {
	repo repository.Repository
}

func NewClientService(repo repository.Repository) *ClientService {
	return &ClientService{
		repo: repo,
	}
}

func (s *ClientService) Check(link string) (int, error) {
	_, err := s.repo.Check(link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 301, nil
		}
		return 0, err
	}
	return 200, nil
}
