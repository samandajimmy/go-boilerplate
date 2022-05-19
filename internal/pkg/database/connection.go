package database

import (
	"fmt"
	"os"
	"reflect"

	"go-boiler-plate/internal/pkg/msg"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"repo.pegadaian.co.id/ms-pds/modules/pgdlogger"
)

type Db struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Logger   string

	Sqlx *sqlx.DB
}

func NewDb(dbArgs ...Db) *Db {
	db := Db{}

	if len(dbArgs) > 0 {
		db = dbArgs[0]
	}

	// check if dbConfig empty or not
	if reflect.DeepEqual(db, Db{}) {
		db = Db{
			Host:     os.Getenv(`DB_HOST`),
			Port:     os.Getenv(`DB_PORT`),
			Username: os.Getenv(`DB_USER`),
			Password: os.Getenv(`DB_PASS`),
			Name:     os.Getenv(`DB_NAME`),
		}
	}

	postgresUrl := fmt.Sprintf("user=%s host=%s dbname=%s sslmode=disable", db.Username, db.Host, db.Name)

	if db.Password != "" {
		postgresUrl += fmt.Sprintf(" password=%s", db.Password)
	}

	if db.Port != "" {
		postgresUrl += fmt.Sprintf(" port=%s", db.Port)
	}

	sqlx, err := sqlx.Connect("postgres", postgresUrl)

	if err != nil {
		pgdlogger.Make().Fatal(err)
	}

	err = sqlx.Ping()

	if err != nil {
		pgdlogger.Make().Fatal(err)
	}

	db.Sqlx = sqlx

	return &db
}

func (db *Db) Migrate() *migrate.Migrate {
	driver, err := postgres.WithInstance(db.Sqlx.DB, &postgres.Config{})

	if err != nil {
		pgdlogger.Make().Fatal(err)
	}

	migrationPath := "migration/postgres"

	if os.Getenv(`APP_PATH`) != "" {
		migrationPath = os.Getenv(`APP_PATH`) + "/" + migrationPath
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		db.Username, driver)

	if err != nil {
		pgdlogger.Make().Fatal(err)
	}

	if err := migration.Up(); err != nil && err.Error() != msg.ErrMigrateNoChange.Error() {
		pgdlogger.Make().Fatal(err)
	}

	return migration
}
