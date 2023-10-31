package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/bhuvana-chinnadurai/farmaze-backend/config" // Replace with the actual module path
)

func Migrate(dbConfig config.DBConfig, migrationPath string) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	fmt.Println("check conn", connStr)
	m, err := migrate.New(
		"file://"+migrationPath, // Path to your migration files
		connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Database migrated")
}
