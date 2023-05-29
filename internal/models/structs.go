package models

type Data struct {
	Id           int
	Active_link  string `json:"active_link"`
	History_link string `json:"history_link"`
}

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}
