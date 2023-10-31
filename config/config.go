package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type MigrationConfig struct {
	Path string
}

type Config struct {
	DB        DBConfig
	Migration MigrationConfig
}

func LoadConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}

	return &Config{
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Migration: MigrationConfig{
			Path: os.Getenv("MIGRATION_PATH"),
		},
	}, nil
}
