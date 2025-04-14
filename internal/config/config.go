package config

import (
	"backend-todo-list/lib/logger"
	"os"
)

type PostgresCfg struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Postgres PostgresCfg   `yaml:"postgres"`
	Logger   logger.Config `yaml:"logger"`
}

func Parse(path string) Config {
	return Config{
		Postgres: PostgresCfg{
			Host: os.Getenv("POSTGRES_HOST"),
			Port: os.Getenv("POSTGRES_PORT"),
		},
		Logger: logger.Config{
			Mode: os.Getenv("LOGGER_MODE"),
		},
	}
}
