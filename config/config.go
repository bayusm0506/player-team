package config

import "os"

type Config struct {
	APPS *APPConfig
}

type APPConfig struct {
	Port string
}

func GetConfig() *Config {
	return &Config{
		APPS: &APPConfig{
			Port: os.Getenv("APPS_PORT"),
		},
	}
}
