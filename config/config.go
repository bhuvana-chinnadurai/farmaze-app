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
	DB         DBConfig
	Migration  MigrationConfig
	Port       string
	CORSConfig CORSConfig
}

type CORSConfig struct {
	AllowedOrigins     []string
	AllowCredentials   bool
	AllowedHeaders     []string
	AllowedMethods     []string
	OptionsPassthrough bool
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
		CORSConfig: CORSConfig{
			AllowedOrigins:     []string{os.Getenv("CORS_ALLOWED_ORIGINS")},
			AllowCredentials:   true,                                     // Set to your specific needs
			AllowedHeaders:     []string{"*"},                            // Adjust based on your requirements
			AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Adjust based on your requirements
			OptionsPassthrough: true,                                     // Set to your specific needs
		},
	}, nil
}
