package repository

import (
	"fmt"
	"test/internal/models"
	"time"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func ConnectDB(logger *zap.SugaredLogger, config *models.Config) (*sqlx.DB, error) {
	return NewPostgresDB(logger, config)
}

func ConnectRedis(logger *zap.SugaredLogger, cfg *models.Config) (*redis.Client, error) {
	return NewRedisCacheDB(logger, cfg)
}

func NewPostgresDB(logger *zap.SugaredLogger, cfg *models.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logger.Errorf("Error while connect to Postgres: %v", err)
		return nil, err
	}
	return db, nil
}

func NewRedisCacheDB(logger *zap.SugaredLogger, cfg *models.Config) (*redis.Client, error) {
	// redisUri := fmt.Sprintf("%s:%s", cfg.HostRedis, cfg.PortRedis)

	client := redis.NewClient(&redis.Options{
		Addr:        "redis:6379",
		DB:          0,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})

	if _, err := client.Ping().Result(); err != nil {
		logger.Error("error while ping redis")
		return nil, err
	}

	defer client.Close()

	return client, nil
}
