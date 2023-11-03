package config

import (
	"fmt"
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
	Port      string
}

func LoadConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}

	var port string
	parsedPort, err := strconv.Atoi(port)
	if err != nil || parsedPort <= 0 || parsedPort > 65535 {
		fmt.Println("Invalid port specified. Using default port 8080.")
		port = "8080" // Set to default 8080
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
		Port: port,
	}, nil
}
