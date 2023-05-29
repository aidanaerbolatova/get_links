package config

import (
	"io/ioutil"
	"test/internal/models"

	"gopkg.in/yaml.v3"
)

func ParseYaml() (models.Config, error) {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return models.Config{}, err
	}
	var config models.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return models.Config{}, err
	}
	return config, nil
}
