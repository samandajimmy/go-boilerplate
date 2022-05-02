package database

import (
	"database/sql"
	"fmt"
	"go-boiler-plate/internal/app/model"
	"net/http"
	"os"

	"repo.pegadaian.co.id/ms-pds/modules/pgdlogger"

	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo/v4"
)

func GetDBConn() (*sql.DB, *pg.DB) {
	dbHost := os.Getenv(`DB_HOST`)
	dbPort := os.Getenv(`DB_PORT`)
	dbUser := os.Getenv(`DB_USER`)
	dbPass := os.Getenv(`DB_PASS`)
	dbName := os.Getenv(`DB_NAME`)

	connection := fmt.Sprintf("postgres://%s%s@%s%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := sql.Open(`postgres`, connection)

	if err != nil {
		pgdlogger.Make().Debug(err)
	}

	err = dbConn.Ping()

	if err != nil {
		pgdlogger.Make().Debug(err)
		os.Exit(1)
	}

	// go-pg connection initiation
	dbOpt, err := pg.ParseURL(connection)

	if err != nil {
		pgdlogger.Make().Debug(err)
	}

	dbpg := pg.Connect(dbOpt)

	return dbConn, dbpg
}

func Ping(echTx echo.Context) error {
	response := model.Response{}
	response.Status = model.StatusSuccess
	response.Message = "PONG!!"
	response.Data = map[string]interface{}{
		"appSlug":    model.AppSlug,
		"appName":    model.AppName,
		"appVersion": model.AppVersion,
		"appHash":    model.BuildHash,
	}

	return echTx.JSON(http.StatusOK, response)
}
