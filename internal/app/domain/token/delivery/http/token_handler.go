package delivery

import (
	"go-boiler-plate/internal/app/domain/token"
	"go-boiler-plate/internal/app/model"
	"go-boiler-plate/internal/app/payload"

	"github.com/labstack/echo/v4"

	"repo.pegadaian.co.id/ms-pds/modules/pgdutil"
)

var (
	hCtrl = pgdutil.NewHandler(&pgdutil.Handler{})
)

type TokenHandler struct {
	TokenUseCase token.ITokenUsecase
}

// NewTokensHandler represent to register tokens endpoint
func NewTokensHandler(echoGroup model.EchoGroup, tknUseCase token.ITokenUsecase) {
	handler := &TokenHandler{
		TokenUseCase: tknUseCase,
	}

	echoGroup.Token.POST("/create", handler.HCreateToken)
	echoGroup.Token.GET("/get", handler.HGetToken)
	echoGroup.Token.GET("/refresh", handler.HRefreshToken)
}

func (t *TokenHandler) HCreateToken(c echo.Context) error {
	var pl model.AccountToken
	var errors pgdutil.ResponseErrors
	err := hCtrl.Validate(c, &pl)

	if err != nil {
		return hCtrl.ShowResponse(c, nil, err, errors)
	}

	response, err := t.TokenUseCase.UCreateToken(c, pl)

	return hCtrl.ShowResponse(c, response, err, errors)
}

func (t *TokenHandler) HGetToken(c echo.Context) error {
	var pl payload.TokenRequest
	var errors pgdutil.ResponseErrors

	_ = echo.QueryParamsBinder(c).
		String("username", &pl.Username).
		String("password", &pl.Password).
		BindError()

	err := hCtrl.Validate(c, &pl)

	if err != nil {
		return hCtrl.ShowResponse(c, nil, err, errors)
	}

	accToken, err := t.TokenUseCase.UGetToken(c, pl.Username, pl.Password)

	return hCtrl.ShowResponse(c, accToken, err, errors)
}

func (t *TokenHandler) HRefreshToken(c echo.Context) error {
	var pl payload.TokenRequest
	var errors pgdutil.ResponseErrors
	_ = echo.QueryParamsBinder(c).
		String("username", &pl.Username).
		String("password", &pl.Password).
		BindError()

	err := hCtrl.Validate(c, &pl)

	if err != nil {
		return hCtrl.ShowResponse(c, nil, err, errors)
	}

	accToken, err := t.TokenUseCase.URefreshToken(c, pl.Username, pl.Password)

	return hCtrl.ShowResponse(c, accToken, err, errors)
}
