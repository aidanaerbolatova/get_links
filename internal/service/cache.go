package service

import (
	"errors"
	"test/internal/repository"

	"go.uber.org/zap"
)

type CacheService struct {
	repo   repository.RedisCache
	logger *zap.SugaredLogger
}

func NewCacheService(repo repository.RedisCache, logger *zap.SugaredLogger) *CacheService {
	return &CacheService{repo: repo, logger: logger}
}

func (s *CacheService) Add(key, value string) error {
	len, err := s.Len()
	if err != nil {
		s.logger.Errorf("error while get len of cache: %v", err)
		return err
	}
	if len > 1000 {
		s.logger.Errorf("error cache is full")
		return errors.New("len of the cache is full")
	}
	return s.repo.Add(key, value)
}

func (s *CacheService) Get(key string) (string, bool, error) {
	return s.repo.Get(key)
}

func (s *CacheService) Len() (int, error) {
	return s.repo.Len()
}
