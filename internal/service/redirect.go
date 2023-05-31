package service

import (
	"errors"
	"net/http"
	"test/internal/repository"
)

var ErrRedirectPage = errors.New("not found active link")

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
		if errors.Is(err, repository.ErrNoRows) {
			return 301, ErrRedirectPage
		}
		return http.StatusBadRequest, err
	}
	return 200, nil
}
