package service

import (
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
	return s.repo.Add(key, value)
}

func (s *CacheService) Get(key string) (string, bool, error) {
	return s.repo.Get(key)
}
