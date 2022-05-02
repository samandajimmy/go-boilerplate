package main

import (
	_tokenHttpDelivery "go-boiler-plate/internal/app/domain/token/delivery/http"
	"go-boiler-plate/internal/app/model"
	"go-boiler-plate/internal/pkg/database"
	"os"
	"strconv"
	"time"

	"repo.pegadaian.co.id/ms-pds/modules/pgdlogger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var ech *echo.Echo

func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	ech = echo.New()
	ech.Debug = true
	loadEnv()
	pgdlogger.Init("debug")
}

func main() {

	dbConn, dbpg := database.GetDBConn()

	defer dbConn.Close()

	echoGroup := model.EchoGroup{
		Token: ech.Group("/token"),
	}

	contextTimeout, err := strconv.Atoi(os.Getenv(`CONTEXT_TIMEOUT`))

	if err != nil {
		pgdlogger.Make().Debug(err)
	}

	timeoutContext := time.Duration(contextTimeout) * time.Second

	repos := newRepositories(dbConn, dbpg)

	usecases := newUsecases(repos, timeoutContext)

	_tokenHttpDelivery.NewTokensHandler(echoGroup, usecases.TokenUseCase)

	// PING
	ech.GET("/", database.Ping)
	ech.GET("/ping", database.Ping)

	// run refresh all token
	_ = usecases.TokenUseCase.URefreshAllToken()

	ech.Start(":" + os.Getenv(`PORT`))

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
