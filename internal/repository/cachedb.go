package repository

import (
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type RedisCache struct {
	redis  *redis.Client
	logger *zap.SugaredLogger
}

func NewRedisCache(redis *redis.Client, logger *zap.SugaredLogger) *RedisCache {
	return &RedisCache{redis: redis, logger: logger}
}

func (r *RedisCache) Add(key, value string) error {
	err := r.redis.Set(key, value, 0).Err()
	if err != nil {
		r.logger.Errorf("error while add key, value to redis: %v", err)
		return err
	}

	err = r.redis.Expire(key, time.Duration(24)*time.Hour).Err()
	if err != nil {
		r.logger.Errorf("error while add ttl: %v", err)
		return err
	}

	return nil
}

func (r *RedisCache) Get(key string) (string, bool, error) {
	value, err := r.redis.Get(key).Result()
	if err != nil {
		r.logger.Errorf("error while get value from redis: %v", err)
		return "", false, err
	}
	return value, true, nil
}
