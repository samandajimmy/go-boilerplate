package database

import (
	"database/sql"
	"fmt"
	"os"

	"repo.pegadaian.co.id/ms-pds/modules/pgdlogger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DbConnection() *sqlx.DB {
	dbHost := os.Getenv(`DB_HOST`)
	dbPort := os.Getenv(`DB_PORT`)
	dbUser := os.Getenv(`DB_USER`)
	dbPass := os.Getenv(`DB_PASS`)
	dbName := os.Getenv(`DB_NAME`)

	postgresUrl := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	sqlx, err := sqlx.Connect("postgres", postgresUrl)

	if err != nil {
		pgdlogger.Make().Fatal(err)
	}

	err = sqlx.Ping()

	if err != nil {
		pgdlogger.Make().Fatal(err)
	}

	return sqlx
}

func DbMigration(sql *sql.DB) *migrate.Migrate {
	driver, err := postgres.WithInstance(sql, &postgres.Config{})

	if err != nil {
		pgdlogger.Make().Debug(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://migration/postgres",
		os.Getenv(`DB_USER`), driver)

	if err != nil {
		pgdlogger.Make().Debug(err)
	}

	if err := migration.Up(); err != nil {
		pgdlogger.Make().Debug(err)
	}

	return migration
}
