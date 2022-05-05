package main

import (
	"net/http"
	"os"
	"time"

	_tokenHttpDelivery "go-boiler-plate/internal/app/domain/token/delivery/http"

	"go-boiler-plate/internal/app/middleware"
	"go-boiler-plate/internal/app/model"
	"go-boiler-plate/internal/pkg/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"repo.pegadaian.co.id/ms-pds/modules/pgdlogger"
)

var ech *echo.Echo

func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	ech = echo.New()
	ech.Debug = true
	loadEnv()
	pgdlogger.Init(os.Getenv(`APP_LOG_LEVEL`))
}

func main() {
	sqlx := database.DbConnection()
	migrate := database.DbMigration(sqlx.DB)

	defer sqlx.Close()
	defer migrate.Close()

	echoGroup := model.EchoGroup{
		Api:   ech.Group("/api"),
		Token: ech.Group("/token"),
	}

	// load all middlewares
	middleware.InitMiddleware(ech, echoGroup)

	repos := newRepositories(sqlx)

	usecases := newUsecases(repos)

	_tokenHttpDelivery.NewTokensHandler(echoGroup, usecases.ITokenUseCase)

	// PING
	ech.GET("/", pingHandler)
	ech.GET("/ping", pingHandler)

	// run refresh all token
	_ = usecases.ITokenUseCase.URefreshAllToken()

	ech.Start(":" + os.Getenv(`APP_PORT`))

}

func loadEnv() {
	// check .env file existence
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return
	}

	err := godotenv.Load()

	if err != nil {
		pgdlogger.Make().Fatal("Error loading .env file")
	}
}

func pingHandler(echTx echo.Context) error {
	response := map[string]interface{}{
		"status":  model.StatusSuccess,
		"message": "PONG!!",
		"data": map[string]interface{}{
			"appSlug":    AppSlug,
			"appName":    AppName,
			"appVersion": AppVersion,
			"appHash":    BuildHash,
		},
	}

	return echTx.JSON(http.StatusOK, response)
}
