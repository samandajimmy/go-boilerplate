package token

import (
	"go-boiler-plate/internal/app/model"

	"github.com/labstack/echo"
)

type UseCase interface {
	UCreateToken(c echo.Context, accToken *model.AccountToken) error
	UGetToken(c echo.Context, username string, password string) (*model.AccountToken, error)
	URefreshToken(c echo.Context, username string, password string) (*model.AccountToken, error)
	URefreshAllToken() error
}
