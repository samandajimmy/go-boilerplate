package test

import (
	"os"

	"go-boiler-plate/internal/pkg/database"

	"github.com/golang-migrate/migrate/v4"
)

func NewTestDb() (*database.Db, *migrate.Migrate) {
	dbConfig := database.Db{
		Host:     os.Getenv(`DB_TEST_HOST`),
		Port:     os.Getenv(`DB_TEST_PORT`),
		Username: os.Getenv(`DB_TEST_USER`),
		Password: os.Getenv(`DB_TEST_PASS`),
		Name:     os.Getenv(`DB_TEST_NAME`),
		Logger:   os.Getenv(`DB_TEST_LOGGER`),
	}

	db := database.NewDb(dbConfig)
	migrator := db.Migrate()

	return db, migrator
}
