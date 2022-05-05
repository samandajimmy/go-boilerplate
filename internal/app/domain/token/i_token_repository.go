package token

import (
	"go-boiler-plate/internal/app/model"

	"github.com/labstack/echo/v4"
)

type ITokenRepository interface {
	RCreate(c echo.Context, accToken *model.AccountToken) error
	RGetByUsername(c echo.Context, accToken *model.AccountToken) error
	RUpdateToken(c echo.Context, accToken *model.AccountToken) error
	RUpdateAllAccountTokenExpiry() error
}
