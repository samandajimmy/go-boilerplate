package token

import (
	"go-boiler-plate/internal/app/model"
	"go-boiler-plate/internal/app/payload"

	"github.com/labstack/echo/v4"
)

type ITokenUsecase interface {
	UCreateToken(c echo.Context, accToken model.AccountToken) (payload.TokenResponse, error)
	UGetToken(c echo.Context, username string, password string) (payload.TokenResponse, error)
	URefreshToken(c echo.Context, username string, password string) (payload.TokenResponse, error)
	URefreshAllToken() error
}
