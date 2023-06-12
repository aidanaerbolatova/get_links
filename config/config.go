package config

import (
	"test/internal/models"

	"github.com/spf13/viper"
)

func ReadConfig() (*models.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &models.Config{
		Host:      viper.GetString("postgres.host"),
		Port:      viper.GetString("postgres.port"),
		Username:  viper.GetString("postgres.username"),
		Password:  viper.GetString("postgres.password"),
		DBName:    viper.GetString("postgres.dbname"),
		SSLMode:   viper.GetString("postgres.sslmode"),
		HostRedis: viper.GetString("redis.host"),
		PortRedis: viper.GetString("redis.port"),
	}, nil
}
