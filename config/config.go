package config

import (
	"io/ioutil"
	"test/internal/models"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func ParseYaml() (*models.Config, error) {
	yamlFile, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		return &models.Config{}, err
	}
	var config models.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return &models.Config{}, err
	}
	return &config, nil
}

func ReadConfig() (*models.Config, error) {
	viper.SetConfigFile("config.yml")
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
