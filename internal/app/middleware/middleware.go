package middleware

import (
	"encoding/json"
	"net/url"
	"os"

	"go-boiler-plate/cmd/router"
	cmdutil "go-boiler-plate/cmd/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gopkg.in/go-playground/validator.v9"

	"github.com/samandajimmy/pgdlogger"
	"github.com/samandajimmy/pgdutil"
)

type CustomValidator struct {
	Validator *validator.Validate
}

type customMiddleware struct {
	e *echo.Echo
}

var (
	echGroup  cmdutil.EchoGroup
	jwtMiddFn echo.MiddlewareFunc
)

// InitMiddleware to generate all middleware that app need
func InitMiddleware(router router.Router) {
	jwtMiddFn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv(`APP_JWT_SECRET`)),
	})

	cm := &customMiddleware{router.Echo}
	echGroup = router.EchoGroup
	router.Echo.Use(middleware.RequestIDWithConfig(middleware.DefaultRequestIDConfig))

	router.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			pgdlogger.SetRequestId(c.Response().Header().Get(echo.HeaderXRequestID))
			return next(c)
		}
	})

	cm.customBodyDump()

	router.Echo.Use(middleware.Recover())
	cm.cors()
	cm.basicAuth()
	cm.jwtAuth()
	cv := CustomValidator{}
	cv.CustomValidation()
	cm.e.Validator = &cv
}

func (cm *customMiddleware) customBodyDump() {
	cm.e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, req, resp []byte) {
			bodyParser(c, &req)
			reqBody := c.Request()
			reqMap, respMap := map[string]interface{}{}, map[string]interface{}{}
			_ = json.Unmarshal(req, &reqMap)
			_ = json.Unmarshal(resp, &respMap)

			pgdlogger.MakeWithoutReportCaller(reqMap).Info("Request payload for endpoint " + reqBody.Method + " " + reqBody.URL.Path)
			pgdlogger.MakeWithoutReportCaller(respMap).Info("Response payload for endpoint " + reqBody.Method + " " + reqBody.URL.Path)
		},
	}))
}

func (cm customMiddleware) cors() {
	cm.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"Access-Control-Allow-Origin"},
		AllowMethods: []string{"*"},
	}))
}

func (cm customMiddleware) basicAuth() {
	echGroup.Token.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == os.Getenv(`APP_BASIC_USERNAME`) && password == os.Getenv(`APP_BASIC_PASSWORD`) {
			return true, nil
		}

		return false, nil
	}))
}

func (cm customMiddleware) jwtAuth() {
	echGroup.Api.Use(jwtMiddFn)
}

func (cv *CustomValidator) CustomValidation() {
	validator := validator.New()

	for key, fn := range pgdutil.WrapCustomValidatorFunc {
		_ = validator.RegisterValidation(key, fn)
	}

	cv.Validator = validator
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func bodyParser(c echo.Context, pl *[]byte) {
	if string(*pl) == "" {
		rawQuery := c.Request().URL.RawQuery
		m, err := url.ParseQuery(rawQuery)

		if err != nil {
			pgdlogger.Make().Fatal(err)
		}

		*pl, err = json.Marshal(m)

		if err != nil {
			pgdlogger.Make().Fatal(err)
		}
	}
}
