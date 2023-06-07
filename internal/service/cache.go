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
