package models

type Data struct {
	Id           int
	Active_link  string `json:"active_link"`
	History_link string `json:"history_link"`
}

type Config struct {
	Host           string `yml:"host"`
	Port           string `yml:"port"`
	Username       string `yml:"username"`
	Password       string `yml:"password"`
	DBName         string `yml:"dbname"`
	SSLMode        string `yml:"sslmode"`
	RequestTimeout int    `yml:"requestTimeout"`
	HostRedis      string `yml:"host_redis"`
	PortRedis      string `yml:"port_redis"`
}
