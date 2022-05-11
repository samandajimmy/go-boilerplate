package main

import (
	"net/http"
	"os"
	"time"

	"go-boiler-plate/cmd/router"
	cmdutil "go-boiler-plate/cmd/util"
	"go-boiler-plate/docs"
	"go-boiler-plate/internal/app/middleware"
	"go-boiler-plate/internal/pkg/database"
	"go-boiler-plate/internal/pkg/msg"

	"github.com/labstack/echo/v4"

	"repo.pegadaian.co.id/ms-pds/modules/pgdlogger"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	cmdutil.LoadEnv()
	pgdlogger.Init(os.Getenv(`APP_LOG_LEVEL`))
}

func main() {
	db := database.NewDb()
	migration := db.Migrate()

	defer db.Sqlx.Close()
	defer migration.Close()

	router := router.NewRoutes(db)

	// load all middlewares
	middleware.InitMiddleware(router)
	// load all handlers
	router.LoadHandlers()

	// PING
	router.Echo.GET("/", pingHandler)
	router.Echo.GET("/ping", pingHandler)

	// run refresh all token
	_ = router.Usecases.ITokenUsecase.URefreshAllToken()

	router.Echo.GET("/swagger/*", documentation())
	router.Echo.Start(":" + os.Getenv(`APP_PORT`))

}

func pingHandler(echTx echo.Context) error {
	response := map[string]interface{}{
		"status":  msg.StatusSuccess,
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

// @termsOfService  http://swagger.io/terms/

// @contact.name   Digital Deparment Pegadaian
// @contact.url    digital@pegadaian.co.id
// @contact.email  digital@pegadaian.co.id

// @license.name  PT. Pegadaian
// @license.url   www.pegadaian.co.id
func documentation() echo.HandlerFunc {

	docs.SwaggerInfo.Title = AppName
	docs.SwaggerInfo.Description = AppDescription
	docs.SwaggerInfo.Version = AppVersion
	docs.SwaggerInfo.Host = os.Getenv("DOC_HOST") + ":" + os.Getenv("DOC_PORT")

	return echoSwagger.WrapHandler
}
